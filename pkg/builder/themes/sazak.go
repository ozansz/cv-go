package themes

const (
	sazakBaseThemeCSS       = "https://sazak.io/static/master.css"
	sazakStyleCustomization = `html, body {
	/* width: 70vw; */
	width: 90vw;
	margin: 0.5vh auto;
}

h1 {
	margin: 0;
	margin-bottom: 0.2em;

	margin-left: -1em;
}

h2, h3, h4, h5, h6 {
	margin: 0;
	margin-bottom: 0.4em;
	font-weight: 400;
}

p {
	margin: 0;
}

.header {
	width: 100%;
	display: flex;
	flex-direction: column;
	margin-top: 2em;
}

.contact {
	display: flex;
	flex-direction: row;
	flex-wrap: wrap;
	justify-content: center;

	margin: 0.8em 0;
}

.contact-item {
	display: flex;
	flex-direction: row;
	align-items: center;
	font-size: 0.8em;
}

.contact-item:not(:last-child) {
	margin-right: 2em;
}

.contact-text {
	margin-left: 0.5em;
}

/* .contact-icon {
	color: rgb(185, 147, 255);
} */
ion-icon {
	color: rgb(185, 147, 255);
}

.header .tldr {
	font-size: 0.8em;
}

.cv-content {
	column-count: 2;
	column-gap: 1.5em;
	margin-top: 1em;
}

.cv-content a {
	color: #f0f0f0;
}

.top-block {
	margin-bottom: 1em;
}

.block {
	margin-bottom: 1.5em;
}

.block-header-row {
	color: #f0f0f0;
	width: 100%;
	display: flex;
	flex-direction: row;
	align-content: space-between;
}

.block-header-row .left {
	flex: 1;
	display: flex;
	flex-direction: row;
	align-items: center;
}

.block-header-row .right {
	display: flex;
	flex-direction: row;
	text-align: right;
	align-items: center;
}

.block-header-row p {
	margin: 0;
}

.block-header-row ion-icon {
	margin-right: 0.5em;
}

.block-header-row .tinted-block-elem {
	color: rgb(185, 147, 255);
	margin-right: 0.5em;
}

.date-start-item::after {
	content: '-';
	padding: 0 0.5em;
}

.block .block-header-row:nth-child(1) {
	margin-bottom: 0;
	margin-bottom: 0.5em;	
}

.block .block-header-row:not(:nth-child(1)) {
	font-size: 0.8em;
}

.block .content-container {
	margin-top: 0.5em;
}

.block .content {
	font-size: 0.8em;
	/* text-align: left; */
}

.block .content:not(:last-child) {
	margin-bottom: 0.8em;
}

a.project-link {
	font-size: 0.8em;
}

ul.bullet-points {
	font-size: 0.8em;
    margin: 0;
    text-align: left;
    padding-left: 1em;
}

a.markdown-link::after, a.markdown-link::before {
	display: none;
}

a.markdown-link {
    display: inline-flex;
	flex-direction: row;
    align-items: center;
	text-decoration: underline;
}

span.markdown-link-text {
	margin-right: 0.1em;
}

table.lang-skills {
	font-size: 0.8em;
	font-weight: 400;
	width: 100%;
	margin-top: 0.8em;
}

table.lang-skills tr th {
	font-weight: 400;
	text-align: center;
}

td.lang-skills-data {
	text-align: center;
}

td.lang-skill-lang {
	color: #f0f0f0;
}

table.cat-skills {
	font-size: 0.8em;
	font-weight: 400;
	width: 100%;
}

td.cat-skills-category {
	color: rgb(185, 147, 255);
}`
)
