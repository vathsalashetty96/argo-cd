package controller

import (
	"context"
	"time"

	"github.com/vathsalashetty96/gitops-engine/pkg/cache"
	"github.com/vathsalashetty96/gitops-engine/pkg/utils/kube"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/vathsalashetty96/argo-cd/controller/metrics"
	appv1 "github.com/vathsalashetty96/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/vathsalashetty96/argo-cd/pkg/client/listers/application/v1alpha1"
	"github.com/vathsalashetty96/argo-cd/util/argo"
	appstatecache "github.com/vathsalashetty96/argo-cd/util/cache/appstate"
	"github.com/vathsalashetty96/argo-cd/util/db"
)

const (
	secretUpdateInterval = 10 * time.Second
)

type clusterInfoUpdater struct {
	infoSource    metrics.HasClustersInfo
	db            db.ArgoDB
	appLister     v1alpha1.ApplicationNamespaceLister
	cache         *appstatecache.Cache
	clusterFilter func(cluster *appv1.Cluster) bool
}

func NewClusterInfoUpdater(
	infoSource metrics.HasClustersInfo,
	db db.ArgoDB,
	appLister v1alpha1.ApplicationNamespaceLister,
	cache *appstatecache.Cache,
	clusterFilter func(cluster *appv1.Cluster) bool) *clusterInfoUpdater {

	return &clusterInfoUpdater{infoSource, db, appLister, cache, clusterFilter}
}

func (c *clusterInfoUpdater) Run(ctx context.Context) {
	c.updateClusters()
	ticker := time.NewTicker(secretUpdateInterval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break
		case <-ticker.C:
			c.updateClusters()
		}
	}
}

func (c *clusterInfoUpdater) updateClusters() {
	infoByServer := make(map[string]*cache.ClusterInfo)
	clustersInfo := c.infoSource.GetClustersInfo()
	for i := range clustersInfo {
		info := clustersInfo[i]
		infoByServer[info.Server] = &info
	}
	clusters, err := c.db.ListClusters(context.Background())
	if err != nil {
		log.Warnf("Failed to save clusters info: %v", err)
	}
	var clustersFiltered []appv1.Cluster
	if c.clusterFilter == nil {
		clustersFiltered = clusters.Items
	} else {
		for i := range clusters.Items {
			if c.clusterFilter(&clusters.Items[i]) {
				clustersFiltered = append(clustersFiltered, clusters.Items[i])
			}
		}
	}
	_ = kube.RunAllAsync(len(clustersFiltered), func(i int) error {
		cluster := clustersFiltered[i]
		if err := c.updateClusterInfo(cluster, infoByServer[cluster.Server]); err != nil {
			log.Warnf("Failed to save clusters info: %v", err)
		}
		return nil
	})
	log.Debugf("Successfully saved info of %d clusters", len(clustersFiltered))
}

func (c *clusterInfoUpdater) updateClusterInfo(cluster appv1.Cluster, info *cache.ClusterInfo) error {
	apps, err := c.appLister.List(labels.Everything())
	if err != nil {
		return err
	}
	var appCount int64
	for _, a := range apps {
		if err := argo.ValidateDestination(context.Background(), &a.Spec.Destination, c.db); err != nil {
			continue
		}
		if a.Spec.Destination.Server == cluster.Server {
			appCount += 1
		}
	}
	now := metav1.Now()
	clusterInfo := appv1.ClusterInfo{
		ConnectionState:   appv1.ConnectionState{ModifiedAt: &now},
		ApplicationsCount: appCount,
	}
	if info != nil {
		clusterInfo.ServerVersion = info.K8SVersion
		if info.LastCacheSyncTime == nil {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusUnknown
		} else if info.SyncError == nil {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusSuccessful
			syncTime := metav1.NewTime(*info.LastCacheSyncTime)
			clusterInfo.CacheInfo.LastCacheSyncTime = &syncTime
			clusterInfo.CacheInfo.APIsCount = int64(info.APIsCount)
			clusterInfo.CacheInfo.ResourcesCount = int64(info.ResourcesCount)
		} else {
			clusterInfo.ConnectionState.Status = appv1.ConnectionStatusFailed
			clusterInfo.ConnectionState.Message = info.SyncError.Error()
		}
	} else {
		clusterInfo.ConnectionState.Status = appv1.ConnectionStatusUnknown
		if appCount == 0 {
			clusterInfo.ConnectionState.Message = "Cluster has no application and not being monitored."
		}
	}

	return c.cache.SetClusterInfo(cluster.Server, &clusterInfo)
}
