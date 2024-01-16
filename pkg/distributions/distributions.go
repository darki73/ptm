package distributions

import (
	"fmt"
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
)

// Distributions represents a structure that contains a named map of distributions.
type Distributions struct {
	// namedMap is a named map of distributions.
	namedMap map[string]Distribution
	// activeDistribution is the active distribution.
	activeDistribution Distribution
}

// NewDistributions returns a new instance of Distributions.
func NewDistributions(baseImage *bi.Configuration) (*Distributions, error) {
	distributions := &Distributions{
		namedMap: map[string]Distribution{
			"ubuntu": NewUbuntu(),
			"debian": NewDebian(),
		},
		activeDistribution: nil,
	}

	if baseImage == nil || !distributions.IsDistributionSupported(baseImage.GetDistribution()) {
		return nil, fmt.Errorf("the distribution '%s' is not supported", baseImage.GetDistribution())
	}

	distributions.activeDistribution = distributions.GetDistribution(baseImage.GetDistribution()).Initialize(baseImage)

	return distributions, nil
}

// IsDistributionSupported returns true if the distribution is supported, otherwise false.
func (distributions *Distributions) IsDistributionSupported(distribution string) bool {
	_, ok := distributions.namedMap[distribution]
	return ok
}

// GetDistributionByName returns a distribution by name.
func (distributions *Distributions) GetDistributionByName(distribution string) Distribution {
	return distributions.namedMap[distribution]
}

// GetDistribution returns a distribution by name.
func (distributions *Distributions) GetDistribution(distribution string) Distribution {
	if !distributions.IsDistributionSupported(distribution) {
		return nil
	}
	return distributions.GetDistributionByName(distribution)
}

// GetActiveDistribution returns the active distribution.
func (distributions *Distributions) GetActiveDistribution() Distribution {
	return distributions.activeDistribution
}
