package ultis

import (
	"strconv"

	intentv1 "github.com/ntnguyencse/L-KaaS/api/v1"
	// "github.com/ntnguyencse/cluster-api-sdk/kubernetes-client/genericiooptions"
	// "k8s.io/cli-runtime/pkg/genericiooptions"
)

// var _ intentv1.Blueprint

func isEqual(old, new *intentv1.Blueprint) bool {
	// Check name
	if old.Name != new.Name {
		return false
	}
	if old.Namespace != new.Namespace {
		return false
	}
	// Compare Spec
	if len(old.Spec.Blueprints) == len(new.Spec.Blueprints) {
		for idx, item := range old.Spec.Blueprints {
			if item.Name != new.Spec.Blueprints[idx].Name {
				return false
			}
			if item.Revision != new.Spec.Blueprints[idx].Revision {
				return false
			}
			if item.Version != new.Spec.Blueprints[idx].Version {
				return false
			}
			if item.Type != new.Spec.Blueprints[idx].Type {
				return false
			}
		}
	} else {
		return false
	}
	// Compare blueprint values
	for key, element := range old.Spec.Values {
		if new.Spec.Values[key] != element {
			return false
		}
	}
	return true
}
func IsEqual(old, new *intentv1.Cluster) bool {
	if old.Name != new.Name {
		return false
	}
	if old.Namespace != new.Namespace {
		return false
	}
	// Compare spec
	// Infrastructure
	if len(old.Spec.Infrastructure) != len(new.Spec.Infrastructure) {
		return false
	} else {
		// Compare Infrastructure Object
		for idx, item := range new.Spec.Infrastructure {
			if item.Name != new.Spec.Infrastructure[idx].Name {
				return false
			}
			if item.Spec.Name != new.Spec.Infrastructure[idx].Spec.Name {
				return false
			}
			if item.Spec.Revision != new.Spec.Infrastructure[idx].Spec.Revision {
				return false
			}
			if item.Spec.Type != new.Spec.Infrastructure[idx].Spec.Type {
				return false
			}
			if item.Spec.Version != new.Spec.Infrastructure[idx].Spec.Version {
				return false
			}
			if len(item.Override) != len(new.Spec.Infrastructure[idx].Override) {
				return false
			} else {
				for key, element := range item.Override {
					if element != new.Spec.Infrastructure[idx].Override[key] {
						return false
					}
				}
			}
		}
		// Compare Software Object
		for idx, item := range new.Spec.Software {
			if item.Name != new.Spec.Software[idx].Name {
				return false
			}
			if item.Spec.Name != new.Spec.Software[idx].Spec.Name {
				return false
			}
			if item.Spec.Revision != new.Spec.Software[idx].Spec.Revision {
				return false
			}
			if item.Spec.Type != new.Spec.Software[idx].Spec.Type {
				return false
			}
			if item.Spec.Version != new.Spec.Software[idx].Spec.Version {
				return false
			}
			if len(item.Override) != len(new.Spec.Software[idx].Override) {
				return false
			} else {
				for key, element := range item.Override {
					if element != new.Spec.Software[idx].Override[key] {
						return false
					}
				}
			}
		}

	}

	return true

}
func NewRevision(oldRevision string) string {
	if oldRevision != "" {
		return "r1"
	} else {
		var newRevision string = "r"
		newRevisionNumber, _ := strconv.ParseInt(oldRevision[1:], 10, 0)
		newRevisionNumber = newRevisionNumber + 1
		revisionNumber := newRevision + strconv.FormatInt(newRevisionNumber, 10)
		return revisionNumber
	}

}
