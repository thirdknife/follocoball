// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Head(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if title == "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("title = \"Mon app\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templa/components/head.templ`, Line: 7, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta charset=\"utf-8\"><!-- Scale to mobiles --><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><!-- Favicon --><link rel=\"icon\" href=\"/static/favicon.ico\"><!-- PWA --><link rel=\"manifest\" href=\"/static/manifest.json\"><!-- Open Graph --><meta property=\"og:locale\" content=\"en_US\"><meta property=\"og:type\" content=\"website\"><meta property=\"og:title\" content=\"{title}\"><meta property=\"og:site_name\" content=\"Gourmet - Healthy Recipes\"><meta property=\"og:url\" content=\"https://gourmet.quimerch.com\"><meta property=\"og:image\" content=\"https://gourmet.quimerch.com/static/hero.webp\"><meta property=\"og:image:width\" content=\"850\"><meta property=\"og:image:height\" content=\"567\"><meta property=\"og:description\" content=\"Gourmet is a recipe planner\"><!-- SEO --><meta name=\"description\" content=\"Gourmet is a recipe planner\"><meta name=\"keywords\" content=\"recipes, food, cooking\"><meta name=\"author\" content=\"Gourmet\"><!-- json-ld --><script type=\"application/ld+json\">\n\t\t{\n\t\t\t\"@context\": \"https://schema.org\",\n\t\t\t\"@type\": \"WebSite\",\n\t\t\t\"name\": \"Gourmet\",\n\t\t\t\"url\": \"https://gourmet.quimerch.com\",\n\t\t\t\"potentialAction\": {\n\t\t\t\"@type\": \"SearchAction\",\n\t\t\t\"target\": \"https://gourmet.quimerch.com/search?q={search_term_string}\",\n\t\t\t\"query-input\": \"required name=search_term_string\"\n\t\t\t}\n\t\t}\n\t</script><!-- tailwindcss --><link href=\"/static/styles.css\" rel=\"stylesheet\"><link rel=\"stylesheet\" type=\"text/css\" href=\"https://cdn.jsdelivr.net/npm/toastify-js/src/toastify.min.css\"><script src=\"https://unpkg.com/franken-ui@1.1.0/dist/js/core.iife.js\" type=\"module\"></script><script src=\"https://unpkg.com/franken-ui@1.1.0/dist/js/icon.iife.js\" type=\"module\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
