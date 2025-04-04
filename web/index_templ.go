// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/TecharoHQ/anubis"
	"github.com/TecharoHQ/anubis/xess"
)

func base(title string, body templ.Component) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 12, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</title><link rel=\"stylesheet\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(xess.URL)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 13, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><style>\n    body,\n    html {\n      height: 100%;\n      display: flex;\n      justify-content: center;\n      align-items: center;\n      margin-left: auto;\n      margin-right: auto;\n    }\n\n    .centered-div {\n      text-align: center;\n    }\n\n    #status {\n      font-variant-numeric: tabular-nums;\n    }\n\n    #progress {\n      display: none;\n      width: min(20rem, 90%);\n      height: 2rem;\n      border-radius: 1rem;\n      overflow: hidden;\n      margin: 1rem 0 2rem;\n      outline-color: #b16286;\n      outline-offset: 2px;\n      outline-style: solid;\n      outline-width: 4px;\n    }\n\n    .bar-inner {\n      background-color: #b16286;\n      height: 100%;\n      width: 0;\n      transition: width 0.25s ease-in;\n    }\n      </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.JSONScript("anubis_version", anubis.Version).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</head><body id=\"top\"><main><center><h1 id=\"title\" class=\".centered-div\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 59, Col: 49}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</h1></center>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = body.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<footer><center><p>Protected by <a href=\"https://github.com/TecharoHQ/anubis\">Anubis</a> from <a href=\"https://techaro.lol\">Techaro</a>. Made with &lt;3 in Canada.</p></center></footer></main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func index() templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<div class=\"centered-div\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = logo().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<p id=\"status\">Loading...</p><script async type=\"module\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/js/main.mjs?cacheBuster=" + anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 80, Col: 116}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"></script><div id=\"progress\" role=\"progressbar\" aria-labelledby=\"status\"><div class=\"bar-inner\"></div></div><details><summary>Why am I seeing this?</summary><p>You are seeing this because the administrator of this website has set up <a href=\"https://github.com/TecharoHQ/anubis\">Anubis</a> to protect the server against the scourge of <a href=\"https://thelibre.news/foss-infrastructure-is-under-attack-by-ai-companies/\">AI companies aggressively scraping websites</a>. This can and does cause downtime for the websites, which makes their resources inaccessible for everyone.</p><p>Anubis is a compromise. Anubis uses a <a href=\"https://anubis.techaro.lol/docs/design/why-proof-of-work\">Proof-of-Work</a> scheme in the vein of <a href=\"https://en.wikipedia.org/wiki/Hashcash\">Hashcash</a>, a proposed proof-of-work scheme for reducing email spam. The idea is that at individual scales the additional load is ignorable, but at mass scraper levels it adds up and makes scraping much more expensive.</p><p>Ultimately, this is a hack whose real purpose is to give a \"good enough\" placeholder solution so that more time can be spent on fingerprinting and identifying headless browsers (EG: via how they do font rendering) so that the challenge proof of work page doesn't need to be presented to users that are much more likely to be legitimate.</p><p>Please note that Anubis requires the use of modern JavaScript features that plugins like <a href=\"https://jshelter.org/\">JShelter</a> will disable. Please disable JShelter or other such plugins for this domain.</p></details><noscript><p>Sadly, you must enable JavaScript to get past this challenge. This is required because AI companies have changed the social contract around how website hosting works. A no-JS solution is a work-in-progress.</p></noscript><div id=\"testarea\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func errorPage(message string) templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<div class=\"centered-div\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = logo().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 104, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, ".</p><button onClick=\"window.location.reload();\">Try again</button><p><a href=\"/\">Go home</a></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func logo() templ.Component {
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
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<svg width=\"365.25659\" height=\"300\" viewBox=\"0 0 365.25659 488.20352\" version=\"1.1\" id=\"svg5\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:svg=\"http://www.w3.org/2000/svg\"><g id=\"layer1\" transform=\"translate(-15.042449,-13.187951)\"><path style=\"stroke-width:1.1183\" d=\"m 179.14789,13.187951 c 0,0 -17.61545,93.935219 -20.7365,93.676869 -3.12108,-0.25835 -26.95923,-92.602257 -26.95923,-92.602257 0,0 -23.59817,101.460937 -29.89254,107.279877 -6.294415,5.819 -19.10538,6.75596 -36.646018,44.50466 -18.423111,39.64787 -21.203919,66.76789 -21.203919,66.76789 L 228.61498,434.43356 h 14.33689 c -20.86833,-45.85034 -42.89373,-150.31452 -44.96115,-197.24148 50.83851,-45.14373 153.05061,-38.10943 158.38292,-40.57534 5.33234,-2.46589 23.9254,-34.65622 23.9254,-34.65622 0,0 -96.22537,-6.5133 -102.42011,-7.75162 -6.19472,-1.23833 -30.53588,-29.20269 -35.80291,-31.98718 -5.26705,-2.78449 -26.30165,-0.14419 -30.54562,-3.59732 C 207.28643,115.17125 179.1479,13.187951 179.14789,13.187951 Z M 198.17638,132.9722 c 13.09943,0.29227 26.86047,5.49018 35.82477,10.66095 -20.7519,12.44717 -45.86659,11.82844 -62.62017,-0.12014 6.81838,-8.02755 16.60696,-10.76811 26.7954,-10.54081 z M 40.322035,256.87799 c -1.731247,9.81898 -3.255402,18.98067 -4.748386,27.67999 L 176.09879,434.91844 h 30.6177 z m -8.498606,48.40777 c -1.495763,8.05786 -2.95718,17.7604 -4.396735,28.05348 l 94.934956,101.5792 h 30.61551 z m -7.351917,50.43685 c -1.18455,9.65933 -2.304037,19.3648 -3.367991,29.15212 l 46.769638,50.04371 h 30.61551 z m -5.617686,50.38444 c -1.126744,11.1122 -1.899033,19.6711 -2.747686,28.81139 h 29.674136 z m 329.932004,35.52553 c -7.47511,0 -13.45815,1.42828 -17.94951,4.28316 -4.49135,2.82917 -6.73816,6.88018 -6.73816,12.15273 0,4.55238 1.64931,8.03578 4.94715,10.45343 3.32925,2.41765 8.40157,4.17975 15.21711,5.28569 l 5.7946,0.92609 c 4.61698,0.74587 8.04219,1.90393 10.27216,3.47284 2.22997,1.54318 3.34396,3.7935 3.34396,6.75127 0,2.31477 -0.76967,4.25627 -2.30867,5.82517 -1.50758,1.54318 -3.5012,2.64904 -5.98244,3.31776 -2.44983,0.64299 -5.35635,0.9654 -8.71702,0.9654 -7.09822,0 -14.52511,-1.78928 -22.2829,-5.36432 v 8.02465 c 8.13468,2.44336 15.56158,3.66503 22.2829,3.66503 3.26644,0 6.25182,-0.20442 8.95291,-0.61593 2.7011,-0.38581 5.16563,-1.00343 7.3956,-1.85218 2.26138,-0.87447 4.16169,-1.96674 5.70069,-3.27844 1.5704,-1.3117 2.77988,-2.91939 3.62791,-4.82265 0.84801,-1.90326 1.27118,-4.05137 1.27118,-6.4433 0,-4.91247 -1.79035,-8.71827 -5.37088,-11.41885 -3.54911,-2.70057 -8.68363,-4.57846 -15.40495,-5.63297 l -5.74873,-0.9654 c -4.99388,-0.77159 -8.4791,-1.78666 -10.45781,-3.04692 -1.9787,-1.28598 -2.96828,-3.24109 -2.96828,-5.8645 0,-2.18617 0.72192,-4.00046 2.16669,-5.44076 1.47618,-1.46603 3.39313,-2.5061 5.74873,-3.12336 2.35561,-0.61728 5.1189,-0.9261 8.29111,-0.9261 5.8105,0 12.31221,1.40108 19.50465,4.20453 v -7.60091 c -7.22386,-1.95469 -14.08653,-2.93116 -20.588,-2.93116 z m -306.889009,1.04185 -26.854372,57.59875 h 10.036263 l 6.40617,-14.77591 h 31.707597 l 6.406176,14.77591 h 9.894283 L 52.686623,442.67443 Z m 47.866093,0 v 57.59875 h 9.234668 v -48.18497 l 31.187768,48.18497 h 12.81453 v -57.59875 h -9.23468 v 48.18498 l -31.18776,-48.18498 z m 71.092546,0 v 35.95581 c 0,7.48443 2.26129,13.15673 6.78404,17.01469 4.52276,3.83223 11.24437,5.74655 20.16426,5.74655 8.91989,0 15.62485,-1.91432 20.11621,-5.74655 4.52275,-3.85796 6.78403,-9.53026 6.78403,-17.01469 v -35.95581 h -9.56229 v 34.99041 c 0,5.86409 -1.30499,10.23752 -3.91186,13.11812 -2.60686,2.85489 -7.08167,4.28316 -13.42609,4.28316 -2.88954,0 -5.37072,-0.28381 -7.44365,-0.84964 -2.04153,-0.56583 -3.83188,-1.50369 -5.37087,-2.81539 -1.539,-1.33743 -2.68628,-3.13813 -3.44008,-5.40146 -0.75379,-2.26333 -1.12921,-5.04267 -1.12921,-8.33479 v -34.99041 z m 71.70413,0 v 57.59875 h 25.25338 c 7.60074,0 13.59825,-1.35106 17.99537,-4.05163 4.42854,-2.70058 6.64425,-6.72439 6.64425,-12.07409 0,-4.038 -1.33608,-7.33057 -4.00576,-9.87682 -2.63829,-2.57198 -6.23346,-4.25687 -10.78763,-5.05418 7.91484,-1.08021 11.87097,-5.18125 11.87097,-12.3056 0,-4.78387 -2.04137,-8.34666 -6.12442,-10.68715 -4.08304,-2.36621 -9.59517,-3.54928 -16.53635,-3.54928 z m 66.19346,0 v 57.59875 h 9.51644 v -57.59875 z m -56.67703,6.40399 h 14.0879 c 9.20255,0 13.80396,2.89296 13.80396,8.67989 0,2.77773 -1.08508,4.92584 -3.25224,6.4433 -2.13574,1.49174 -5.65206,2.23659 -10.55172,2.23659 h -14.0879 z m -194.808323,1.27337 12.956502,28.66505 H 34.359251 Z m 194.808323,22.41395 h 15.26517 c 5.43359,0 9.29642,0.91285 11.5892,2.73895 2.3242,1.80038 3.48594,4.39812 3.48594,7.79312 0,3.67792 -1.28615,6.36646 -3.86162,8.06396 -2.54404,1.67178 -6.28246,2.50743 -11.21352,2.50743 h -15.26517 z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func bench() templ.Component {
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
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<div style=\"height:20rem;display:flex\"><table style=\"margin-top:1rem;display:grid;grid-template:auto 1fr/auto auto;gap:0 0.5rem\"><thead style=\"border-bottom:1px solid black;padding:0.25rem 0;display:grid;grid-template:1fr/subgrid;grid-column:1/-1\"><tr id=\"table-header\" style=\"display:contents\"><th style=\"width:4.5rem\">Time</th><th style=\"width:4rem\">Iters</th></tr><tr id=\"table-header-compare\" style=\"display:none\"><th style=\"width:4.5rem\">Time A</th><th style=\"width:4rem\">Iters A</th><th style=\"width:4.5rem\">Time B</th><th style=\"width:4rem\">Iters B</th></tr></thead> <tbody id=\"results\" style=\"padding-top:0.25rem;display:grid;grid-template-columns:subgrid;grid-auto-rows:min-content;grid-column:1/-1;row-gap:0.25rem;overflow-y:auto;font-variant-numeric:tabular-nums\"></tbody></table><div class=\"centered-div\"><img id=\"image\" style=\"width:100%;max-width:256px;\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/img/pensive.webp?cacheBuster=" +
			anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 154, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\"><p id=\"status\" style=\"max-width:256px\">Loading...</p><script async type=\"module\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/js/bench.mjs?cacheBuster=" + anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 157, Col: 118}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\"></script><div id=\"sparkline\"></div><noscript><p>Running the benchmark tool requires JavaScript to be enabled.</p></noscript></div></div><form id=\"controls\" style=\"position:fixed;top:0.5rem;right:0.5rem\"><div style=\"display:flex;justify-content:end\"><label for=\"difficulty-input\" style=\"margin-right:0.5rem\">Difficulty:</label> <input id=\"difficulty-input\" type=\"number\" name=\"difficulty\" style=\"width:3rem\"></div><div style=\"margin-top:0.25rem;display:flex;justify-content:end\"><label for=\"algorithm-select\" style=\"margin-right:0.5rem\">Algorithm:</label> <select id=\"algorithm-select\" name=\"algorithm\"></select></div><div style=\"margin-top:0.25rem;display:flex;justify-content:end\"><label for=\"compare-select\" style=\"margin-right:0.5rem\">Compare:</label> <select id=\"compare-select\" name=\"compare\"><option value=\"NONE\">-</option></select></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
