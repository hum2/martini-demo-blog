package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
)

type Profile struct {
	Name  string
	Skill []string
}

type AboutViewModel struct {
	Title   string
	Profile Profile
}

func AboutRender(r render.Render, params martini.Params) {

	skill := []string{"TypeScript", "Sass/Compass", "Go"}

	profile := Profile{
		params["name"],
		skill,
	}

	viewModel := AboutViewModel{
		"About me",
		profile,
	}

	r.HTML(200, "about", viewModel)
}