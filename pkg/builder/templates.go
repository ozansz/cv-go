package builder

const (
	cvTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
	<link rel="stylesheet" href="{{.StyleFile}}">

	<script type="module" src="https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.esm.js"></script>
	<script nomodule src="https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.js"></script>

	<title>cv-go</title>
</head>
<body>
	<div class="header">
		<h1 class="name">{{.Conf.Header.Name}}</h1>
		<p class="title">{{.Conf.Header.Title}}</p>
		<div class="contact">
			{{range .Conf.Header.Contacts}}
			<a href="{{.URL}}" class="contact-item">
				<ion-icon class="contact-icon" name="{{.Icon}}"></ion-icon>
				<div class="contact-text">{{.Text}}</div>
			</a>
			{{end}}
		</div>
		<p class="tldr">{{.Conf.Header.Description}}</p>
	</div>
	<div class="cv-content">
		<div class="top-block">
			<h2 class="subtitle">Experience</h2>
			{{range .Conf.Experiences}}
			<div class="block">
				<div class="block-header-row">
					<div class="left">{{.Title}}</div>
					<div class="right">{{.Company}}</div>
				</div>
				<div class="block-header-row">
					<div class="left">
						<ion-icon name="calendar-outline"></ion-icon>
						<p class="date-start-item">{{.StartDate}}</p>
						<p>{{if .EndDate}}{{.EndDate}}{{else}}Current{{end}}</p>
					</div>
					<div class="right">
						<ion-icon name="location-outline"></ion-icon>
						<p>{{.Location}}</p>
					</div>
				</div>
				<p class="content">{{.Description}}</p>
			</div>
			{{end}}
		</div>
		<div class="top-block">
			<h2 class="subtitle">Projects</h2>
			{{range .Conf.Projects}}
			<div class="block">
				<div class="block-header-row">
					<div class="left">{{.Name}}</div>
					{{if .GithubRepo}}
					<a class="right" href="https://github.com/{{.GithubRepo}}">
						<ion-icon name="logo-github"></ion-icon>
						<div>{{.GithubRepo}}</div>
					</a>
					{{end}}
				</div>
				<div class="block-header-row">
					<div class="left">
						<ion-icon name="calendar-outline"></ion-icon>
						<p class="date-start-item">{{.StartDate}}</p>
						<p>{{if .EndDate}}{{.EndDate}}{{else}}Current{{end}}</p>
					</div>
				</div>
				<p class="content">{{.Description}}</p>
			</div>
			{{end}}
		</div>
	</div>
</body>
</html>
	`
)

type CVTemplateData struct {
	Conf      *Config
	StyleFile string
}
