package builder

const (
	cvTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
	{{if .Theme}}<meta name="cv-go-theme" content="{{.Theme}}">{{end}}

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

				<div class="content-container">
					{{range .DescriptionLines}}
					<p class="content">{{.}}</p>
					{{end}}

					<ul class="bullet-points">
						{{range .BulletPoints}}
						<li>{{.}}</li>
						{{end}}
					</ul>
				</div>
			</div>
			{{end}}
		</div>

		<div class="top-block">
			<h2 class="subtitle">Education</h2>
			{{range .Conf.Education}}
			<div class="block">
				<div class="block-header-row">
					<div class="left">{{.Title}}</div>
				</div>

				<div class="block-header-row">
					<div class="left">
						<ion-icon name="business-outline"></ion-icon>
						<p>{{.School}}</p>
					</div>
					{{if .CGPA}}
					<div class="right">
						<p class="tinted-block-elem">CGPA:</p>
						<p>{{.CGPA}}</p>
					</div>
					{{end}}
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

				<div class="content-container">
					{{range .DescriptionLines}}
					<p class="content">{{.}}</p>
					{{end}}

					<ul class="bullet-points">
						{{range .BulletPoints}}
						<li>{{.}}</li>
						{{end}}
					</ul>
				</div>
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
					<a class="right project-link" href="https://github.com/{{.GithubRepo}}">
						<ion-icon name="logo-github"></ion-icon>
						<div>{{.GithubRepo}}</div>
					</a>
					{{else}}
					{{if .Link}}
					<a class="right project-link" href="{{.Link.URL}}">
						<ion-icon name="link-outline"></ion-icon>
						<div>{{.Link.Title}}</div>
					</a>
					{{end}}
					{{end}}
				</div>
				<div class="block-header-row">
					<div class="left">
						<ion-icon name="calendar-outline"></ion-icon>
						<p class="date-start-item">{{.StartDate}}</p>
						<p>{{if .EndDate}}{{.EndDate}}{{else}}Current{{end}}</p>
					</div>
				</div>
				
				<div class="content-container">
					{{range .DescriptionLines}}
					<p class="content">{{.}}</p>
					{{end}}
				</div>
			</div>
			{{end}}
		</div>

		{{if .Conf.Skills}}
		<div class="top-block">
			<h2 class="subtitle">Skills</h2>
			<div class="block">
				<div class="content-container">
					<table class="cat-skills">
					{{range $category, $skills := .Conf.Skills.CategoricalSkills}}
					<tr>	
						<td class="cat-skills-category">{{$category}}</td>
						<td>{{$skills}}</td>
					</tr>
					{{end}}
					</table>

					{{if .Conf.Skills.Languages}}
					<table class="lang-skills">
					<tr>
						<th></th>
						<th>Reading</th> 
						<th>Writing</th>
						<th>Speaking</th>
						<th>Listening</th>
					</tr>
					{{range $lang, $skill := .Conf.Skills.Languages}}
					<tr>
						<td class="lang-skill-lang">{{$lang}}</td>
						{{if .Native}}
						<td class="lang-skills-data" colspan="4">Native</td>
						{{else}}
						<td class="lang-skills-data">{{if $skill.Reading}}{{$skill.Reading}}{{else}} - {{end}}</td>
						<td class="lang-skills-data">{{if $skill.Writing}}{{$skill.Writing}}{{else}} - {{end}}</td>
						<td class="lang-skills-data">{{if $skill.Speaking}}{{$skill.Speaking}}{{else}} - {{end}}</td>
						<td class="lang-skills-data">{{if $skill.Listening}}{{$skill.Listening}}{{else}} - {{end}}</td>
						{{end}}
					</tr>
					{{end}}
					</table>
					{{end}}
				</div>
			</div>
		</div>
		{{end}}
	</div>
</body>
</html>
	`
)

type CVTemplateData struct {
	Conf      *Config
	StyleFile string
	Theme     string
}
