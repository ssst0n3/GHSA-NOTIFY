package ghsa

type QueryRepository struct {
	Repository `graphql:"repository(name: $name, owner: $owner)"`
}

type Repository struct {
	Description         string                                 `json:"description"`
	Name                string                                 `json:"name"`
	Owner               RepositoryOwner                        `json:"owner"`
	SecurityPolicyUrl   string                                 `json:"securityPolicyUrl"`
	VulnerabilityAlerts RepositoryVulnerabilityAlertConnection `json:"vulnerabilityAlerts" graphql:"vulnerabilityAlerts(first: $first)"`
}

type RepositoryOwner struct {
	Url string `json:"url"`
}

type RepositoryVulnerabilityAlertConnection struct {
	TotalCount int                            `json:"totalCount"`
	Nodes      []RepositoryVulnerabilityAlert `json:"nodes"`
}

type RepositoryVulnerabilityAlert struct {
	SecurityAdvisory SecurityAdvisory `json:"securityAdvisory"`
}

// SecurityVulnerabilities
// https://docs.github.com/cn/graphql/reference/queries#securityvulnerabilityconnection
type SecurityVulnerabilities struct {
	SecurityVulnerabilityConnection `graphql:"securityVulnerabilities(package: $package, first: $first, orderBy: {field: UPDATED_AT, direction: DESC})"`
}

// SecurityVulnerabilityConnection
// https://docs.github.com/cn/graphql/reference/objects#securityvulnerabilityconnection
type SecurityVulnerabilityConnection struct {
	TotalCount int                     `json:"totalCount"`
	Nodes      []SecurityVulnerability `json:"nodes"`
}

// SecurityVulnerability
// https://docs.github.com/en/graphql/reference/objects#securityvulnerability
type SecurityVulnerability struct {
	Advisory SecurityAdvisory        `json:"advisory"`
	Package  SecurityAdvisoryPackage `json:"package"`
}

// SecurityAdvisory
// https://docs.github.com/en/graphql/reference/objects#securityadvisory
type SecurityAdvisory struct {
	CVSS                   CVSS                        `json:"cvss"`
	GHSAId                 string                      `json:"ghsaid"`
	NotificationsPermalink string                      `json:"notificationsPermalink"`
	Origin                 string                      `json:"origin"`
	Permalink              string                      `json:"permalink"`
	Description            string                      `json:"description"`
	References             []SecurityAdvisoryReference `json:"references"`
	Summary                string                      `json:"summary"`
}

// SecurityAdvisoryPackage
// https://docs.github.com/en/graphql/reference/objects#securityadvisorypackage
type SecurityAdvisoryPackage struct {
	Name string
}

// CVSS
// https://docs.github.com/en/graphql/reference/objects#cvss
type CVSS struct {
	Score        float32
	VectorString string
}

// SecurityAdvisoryReference
// https://docs.github.com/en/graphql/reference/objects#securityadvisoryreference
type SecurityAdvisoryReference struct {
	Url string `json:"url"`
}
