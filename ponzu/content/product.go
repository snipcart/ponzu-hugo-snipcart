package content

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Product struct {
	item.Item

	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

// MarshalEditor writes a buffer of html to edit a Product within the CMS
// and implements editor.Editable
func (p *Product) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Product field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Price", p, map[string]string{
				"label":       "Price",
				"type":        "text",
				"placeholder": "Enter the Price here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", p, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Image", p, map[string]string{
				"label":       "Image",
				"placeholder": "Upload the Image here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Product editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Product"] = func() interface{} { return new(Product) }
}

// String defines how a Product is printed. Update it using more descriptive
// fields from the Product struct type
func (p *Product) String() string {
	return fmt.Sprintf("Product: %s", p.UUID)
}

func (p *Product) AfterAdminCreate(res http.ResponseWriter, req *http.Request) error {
	sendWebHook()
	return nil
}

func (p *Product) AfterAdminUpdate(res http.ResponseWriter, req *http.Request) error {
	sendWebHook()
	return nil
}

func (p *Product) AfterAdminDelete(res http.ResponseWriter, req *http.Request) error {
	sendWebHook()
	return nil
}

func sendWebHook() {
	url := os.Getenv("NETLIFY_BUILD_HOOK_URL")
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Webhook called successfully at " + url + " with result " + resp.Status)
	}
}
