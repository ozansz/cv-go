package themes

const (
	sazakLightBaseThemeCSS           = "https://sazak.io/static/master.css"
	sazakLightBaseStyleCustomization = sazakStyleCustomization
	sazakLightStyleCustomization     = `
html, body {
	background: none;
}

h1, h2, h3, h4, h5, h6, p, span, div, li, table, th, td {
	color: #000 !important;
}

h1::before, h2::before, td.cat-skills-category, ion-icon, p.tinted-block-elem {
	color: rgb(0, 132, 255) !important;
}

a.markdown-link, a.project-link {
	text-decoration-color: rgb(0, 132, 255) !important;
}

div.block-header-row:first-child div.left, td.cat-skills-category, p.tinted-block-elem, div.contact-text {
	font-weight: 500;
}

li::before {
	color: #ec8a00 !important;
	font-weight: 600;
}

div.top-block > hr {
	color: #000 !important;
}
`
)
