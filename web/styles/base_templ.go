// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package styles

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

var baseCSSHandle = templ.NewOnceHandle()

func Base() templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style type=\"text/css\">\n@layer reset, base, components;\n\n@layer reset {\n\t*,\n\t*::before,\n\t*::after {\n\t\tbox-sizing: border-box;\n\t\tborder: 0;\n\t\tborder-style: solid;\n\t}\n}\n\n@layer base {\n\t:root {\n\t\t--bg-color: #ffffff;\n\t\t--text-color: #000000;\n\t}\n\n\t:root[data-theme=\"dark\"] {\n\t\t--bg-color: #1b1b1f;\n\t\t--text-color: #f9fafb;\n\t}\n\n\t@media (prefers-color-scheme: dark) {\n\t\t:root {\n\t\t\t--bg-color: #1b1b1f;\n\t\t\t--text-color: #f9fafb;\n\t\t}\n\t}\n\n\tbody {\n\t\tfont-family: ui-sans-serif, system-ui, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji';\n\t\tcolor: var(--text-color);\n\t\tbackground-color: var(--bg-color);\n\t}\n\n\t.dark {\n\t\tcolor: var(--text-color);\n\t\tbackground-color: var(--bg-color);\n\t}\n\n\t.inline {\n\t\tdisplay: flex;\n\t\tflex-wrap: wrap;\n\t\tgap: 0.75rem;\n\t}\n}\n</style>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = baseCSSHandle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
