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
