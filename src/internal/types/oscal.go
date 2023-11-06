package types

type Link struct {
	Rel       string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
	Href      string `json:"href" yaml:"href"`
}
type Metadata struct {
	Version            string             `json:"version" yaml:"version"`
	DocumentIds        []DocumentId       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Published          string             `json:"published,omitempty" yaml:"published,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Roles              []Role             `json:"roles,omitempty" yaml:"roles,omitempty"`
	Parties            []Party            `json:"parties,omitempty" yaml:"parties,omitempty"`
	LastModified       string             `json:"last-modified" yaml:"last-modified"`
	OscalVersion       string             `json:"oscal-version" yaml:"oscal-version"`
	Title              string             `json:"title" yaml:"title"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Locations          []Location         `json:"locations,omitempty" yaml:"locations,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          []Revision         `json:"revisions,omitempty" yaml:"revisions,omitempty"`
}
type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}
type Statement struct {
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalComponentDocument struct {
	ComponentDefinition ComponentDefinition `json:"component-definition" yaml:"component-definition"`
}
type ComponentDefinition struct {
	UUID                       string                      `json:"uuid" yaml:"uuid"`
	Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
	ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}
type DocumentId struct {
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Identifier string `json:"identifier" yaml:"identifier"`
}
type Address struct {
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
}
type ResponsibleRole struct {
	RoleId     string     `json:"role-id" yaml:"role-id"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}
type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}
type PortRange struct {
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
}
type Citation struct {
	Text  string     `json:"text" yaml:"text"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
}
type Protocol struct {
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string      `json:"name" yaml:"name"`
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
}
type Resources struct {
	Remarks     string       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID        string       `json:"uuid" yaml:"uuid"`
	Title       string       `json:"title,omitempty" yaml:"title,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds []DocumentId `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Citation    []Citation   `json:"citation,omitempty" yaml:"citation,omitempty"`
	Rlinks      []Rlinks     `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Base64      []Base64     `json:"base64,omitempty" yaml:"base64,omitempty"`
	Props       []Property   `json:"props,omitempty" yaml:"props,omitempty"`
}
type Property struct {
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Value   string `json:"value" yaml:"value"`
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Name    string `json:"name" yaml:"name"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ExternalIds struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}
type SetParameter struct {
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ParamId string   `json:"param-id" yaml:"param-id"`
	Values  []string `json:"values" yaml:"values"`
}
type DefinedComponent struct {
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Title                  string                  `json:"title" yaml:"title"`
	Description            string                  `json:"description" yaml:"description"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              []Protocol              `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Type                   string                  `json:"type" yaml:"type"`
	Purpose                string                  `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}
type Capability struct {
	IncorporatesComponents []IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Name                   string                  `json:"name" yaml:"name"`
	Description            string                  `json:"description" yaml:"description"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
}
type Rlinks struct {
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
}
type ImplementedRequirement struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Description      string            `json:"description" yaml:"description"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}
type ControlImplementation struct {
	Source                  string                   `json:"source" yaml:"source"`
	Description             string                   `json:"description" yaml:"description"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	UUID                    string                   `json:"uuid" yaml:"uuid"`
}
type TelephoneNumber struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Number string `json:"number" yaml:"number"`
}
type Location struct {
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Address          Address           `json:"address" yaml:"address"`
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
}
type ResponsibleParty struct {
	RoleId     string     `json:"role-id" yaml:"role-id"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}
type Party struct {
	Type                  string            `json:"type" yaml:"type"`
	Name                  string            `json:"name,omitempty" yaml:"name,omitempty"`
	ShortName             string            `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	ExternalIds           []ExternalIds     `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Props                 []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Addresses             []Address         `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	LocationUuids         []string          `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	UUID                  string            `json:"uuid" yaml:"uuid"`
	EmailAddresses        []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers      []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	MemberOfOrganizations []string          `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Remarks               string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Links                 []Link            `json:"links,omitempty" yaml:"links,omitempty"`
}
type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}
type BackMatter struct {
	Resources []Resources `json:"resources,omitempty" yaml:"resources,omitempty"`
}
type Role struct {
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ID          string     `json:"id" yaml:"id"`
	Title       string     `json:"title" yaml:"title"`
	ShortName   string     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
}
type Revision struct {
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Version      string     `json:"version" yaml:"version"`
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
}
type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}
