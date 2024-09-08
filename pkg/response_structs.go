package yadisk

type Disk struct {
	PaidMaxFileSize                 uint          `json:"paid_max_file_size"`
	MaxFileSize                     uint          `json:"max_file_size"`
	TotalSpace                      uint          `json:"total_space"`
	TrashSize                       uint          `json:"trash_size"`
	UsedSpace                       uint          `json:"used_space"`
	IsPaid                          bool          `json:"is_paid"`
	IsIdmManagedFolderAddressAccess bool          `json:"is_idm_managed_folder_address_access"`
	RegTime                         string        `json:"reg_time"`
	SystemFolders                   SystemFolders `json:"system_folders"`
	User                            User          `json:"user"`
	IsIdmManagedPublicAccess        bool          `json:"is_idm_managed_public_access"`
	PaymentFlow                     bool          `json:"payment_flow"`
	UnlimitedAutouploadEnabled      bool          `json:"unlimited_autoupload_enabled"`
	Revision                        uint          `json:"revision"`
}

type SystemFolders struct {
	Attach        string `json:"attach"`
	Applications  string `json:"applications"`
	Calendar      string `json:"calendar"`
	Downloads     string `json:"downloads"`
	Facebook      string `json:"facebook"`
	Social        string `json:"social"`
	Google        string `json:"google"`
	Instagram     string `json:"instagram"`
	MailRu        string `json:"mailru"`
	Messenger     string `json:"messenger"`
	Odnoklassniki string `json:"odnoklassniki"`
	Photostream   string `json:"photostream"`
	Scans         string `json:"scans"`
	Screenshots   string `json:"screenshots"`
	Vkontakte     string `json:"vkontakte"`
}

type User struct {
	Uid         string `json:"uid"`
	DisplayName string `json:"display_name"`
	Login       string `json:"login"`
	Country     string `json:"country"`
	IsChild     bool   `json:"is_child"`
	RegTime     string `json:"reg_time"`
}

type Link struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

type Resource struct {
	AntivirusStatus  map[string]string `json:"antivirus_status"`
	ResourceId       string            `json:"resource_id"`
	Share            ShareInfo         `json:"share"`
	File             string            `json:"file"`
	Size             int               `json:"size"`
	PhotosliceTime   string            `json:"photoslice_time"`
	Embedded         ResourceList      `json:"_embedded"`
	Exif             Exif              `json:"exif"`
	CustomProperties map[string]string `json:"custom_properties"`
	MediaType        string            `json:"media_type"`
	Preview          string            `json:"preview"`
	Type             string            `json:"type"`
	MimeType         string            `json:"mime_type"`
	Revision         uint              `json:"revision"`
	PublicUrl        string            `json:"public_url"`
	Path             string            `json:"path"`
	Md5              string            `json:"md5"`
	PublicKey        string            `json:"public_key"`
	Sha256           string            `json:"sha256"`
	Name             string            `json:"name"`
	Created          string            `json:"created"`
	Sizes            []Preview         `json:"sizes"`
	Modified         string            `json:"modified"`
	CommentIds       CommentIds        `json:"comment_ids"`
}

type ShareInfo struct {
	IsRoot  bool   `json:"is_root"`
	IsOwned bool   `json:"is_owned"`
	Rights  string `json:"rights"`
}

type ResourceList struct {
	Sort   string     `json:"sort"`
	Items  []Resource `json:"items"`
	Limit  uint       `json:"limit"`
	Offset uint       `json:"offset"`
	Path   string     `json:"path"`
	Total  uint       `json:"total"`
}

type Exif struct {
	DateTime     string `json:"date_time"`
	GpsLongitude string `json:"gps_longitude"`
	GpsLatitude  string `json:"gps_latitude"`
}

type Preview struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type CommentIds struct {
	PrivateResource string `json:"private_resource"`
	PublicResource  string `json:"public_resource"`
}

type FilesResourceList struct {
	Items  []Resource `json:"items"`
	Limit  uint       `json:"limit"`
	Offset uint       `json:"offset"`
}

type LastUploadedResourceList struct {
	Items []Resource `json:"items"`
	Limit uint       `json:"limit"`
}

type ResourceUploadLink struct {
	OperationId string `json:"operation_id"`
	Href        string `json:"href"`
	Method      string `json:"method"`
	Templated   bool   `json:"templated"`
}

type PublicResource struct {
	Resource
	Embedded   PublicResourceList    `json:"_embedded"`
	ViewsCount uint                  `json:"views_count"`
	Owner      UserPublicInformation `json:"owner"`
}

type UserPublicInformation struct {
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	Uid         string `json:"uid"`
}

type PublicResourceList struct {
	ResourceList
	PublicKey string           `json:"public_key"`
	Items     []PublicResource `json:"items"`
}
