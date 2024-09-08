package yadisk

// Disk
type DiskGetOpts struct {
	Fields string `url:"fields"`
}

// Resources
type ResourcesGetOpts struct {
	Path        string `url:"path,required"`
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	Offset      uint   `url:"offset"`
	PreviewCrop bool   `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
	Sort        string `url:"sort"`
}

type ResourcesCreateOpts struct {
	Path   string `url:"path,required"`
	Fields string `url:"fields"`
}

type ResourcesUpdateOpts struct {
	Path   string `url:"path,required"`
	Fields string `url:"fields"`
}

type ResourcesUpdateBody struct {
	CustomProperties map[string]string `json:"custom_properties"`
}

type ResourcesDeleteOpts struct {
	Path        string `url:"path,required"`
	Fields      string `url:"fields"`
	ForceAsync  bool   `url:"force_async"`
	Md5         string `url:"md5"`
	Permanently bool   `url:"permanently"`
}

type ResourcesCopyOpts struct {
	Path       string `url:"path,required"`
	From       string `url:"from,required"`
	Fields     string `url:"fields"`
	ForceAsync bool   `url:"force_async"`
	Overwrite  bool   `url:"overwrite"`
}

type ResourcesGetDownloadLinkOpts struct {
	Path   string `url:"path,required"`
	Fields string `url:"fields"`
}

type ResourcesGetFilesOpts struct {
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	MediaType   string `url:"media_type"`
	Offset      uint   `url:"offset"`
	PreviewCrop string `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
	Sort        string `url:"sort"`
}

type ResourcesGetLastUploadedOpts struct {
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	MediaType   string `url:"media_type"`
	Offset      uint   `url:"offset"`
	PreviewCrop string `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
}

type ResourcesMoveOpts struct {
	From       string `url:"from,required"`
	Path       string `url:"path,required"`
	Fields     string `url:"fields"`
	ForceAsync bool   `url:"force_async"`
	Overwrite  bool   `url:"overwrite"`
}

type ResourcesGetPublicOpts struct {
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	Offset      uint   `url:"offset"`
	PreviewCrop string `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
	Type        string `url:"type"`
}

type ResourcesPublishOpts struct {
	Path               string `url:"path,required"`
	Fields             string `url:"fields"`
	AllowAddressAccess bool   `url:"allow_address_access"`
}

type Verbose struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type OrgId Verbose
type Password Verbose
type AvailableUntil Verbose

type SetPublicSettings struct {
	ReadOnly              bool                `json:"read_only"`
	ExternalOrgIdVerbose  OrgId               `json:"external_organization_id_verbose"`
	PasswordVerbose       Password            `json:"password_verbose"`
	AvailableUntil        uint                `json:"available_until"`
	Accesses              []map[string]string `json:"accesses"`
	AvailableUntilVerbose AvailableUntil      `json:"available_until_verbose"`
	Password              string              `json:"password"`
	ExternalOrgId         string              `json:"external_organization_id"`
}

type ResourcesPublishBody struct {
	PublicSettings SetPublicSettings `json:"public_settings"`
}

type ResourcesUnpublishOpts struct {
	Path   string `url:"path,required"`
	Fields string `url:"fields"`
}

type ResourcesGetUploadLinkOpts struct {
	Path      string `url:"path,required"`
	Fields    string `url:"fields"`
	Overwrite bool   `url:"overwrite"`
}

// Public
type PublicGetOpts struct {
	PublicKey   string `url:"public_key,required"`
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	Offset      uint   `url:"offset"`
	Path        string `url:"path"`
	PreviewCrop bool   `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
	Sort        string `url:"sort"`
}

type PublicGetDownloadLinkOpts struct {
	PublicKey string `url:"public_key,required"`
	Fields    string `url:"fields"`
	Path      string `url:"path"`
}

type PublicSaveToDiskOpts struct {
	PublicKey  string `url:"public_key,required"`
	Fields     string `url:"fields"`
	ForceAsync bool   `url:"force_async"`
	Name       string `url:"name"`
	Path       string `url:"path"`
	SavePath   string `url:"save_path"`
}

// Trash
type TrashGetOpts struct {
	Path        string `url:"path,required"`
	Fields      string `url:"fields"`
	Limit       uint   `url:"limit,omitempty"`
	Offset      uint   `url:"offset"`
	PreviewCrop bool   `url:"preview_crop"`
	PreviewSize string `url:"preview_size"`
	Sort        string `url:"sort"`
}

type TrashClearOpts struct {
	Fields     string `url:"fields"`
	ForceAsync bool   `url:"force_async"`
	Path       string `url:"path"`
}

type TrashRestoreOpts struct {
	Path       string `url:"path,required"`
	Fields     string `url:"fields"`
	ForceAsync bool   `url:"force_async"`
	Name       string `url:"name"`
	Overwrite  bool   `url:"overwrite"`
}
