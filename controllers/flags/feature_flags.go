package flags

import "github.com/RHEcosystemAppEng/dbaas-operator/api/v1alpha1"

// Toggle Router decides the current state of a given feature -- whether it is On or Off

type ObsTogglerRouter struct {
	Config       map[string]bool `json:"config,omitempty"`
	toggleConfig *v1alpha1.ToggleConfig
}

func NewObsToggleRouter() *ObsTogglerRouter {
	config := &v1alpha1.ToggleConfig{}
	return &ObsTogglerRouter{toggleConfig: config}
}

func (router ObsTogglerRouter) IsObsConnectionMetricOn() bool {
	return router.Config["obs_connection_all"]
}

func (router ObsTogglerRouter) IsObsConnectionStatusMetricOn() bool {
	return router.Config["obs_connection_status"]
}

func (router ObsTogglerRouter) IsObsRemoteWriteOn() bool {
	return router.Config["obs_connection_remotewrite_all"]
}

type ProviderToggleRouter interface {
}
