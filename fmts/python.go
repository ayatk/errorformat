package fmts

func init() {
	const lang = "python"

	register(&Fmt{
		Name: "pep8",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "Python style guide checker",
		URL:         "https://pypi.python.org/pypi/pep8",
		Language:    lang,
	})

	register(&Fmt{
		Name: "flake8",
		Errorformat: []string{
			`%f:%l:%c: %t%n %m`,
		},
		Description: "Tool for python style guide enforcement",
		URL:         "https://flake8.pycqa.org/",
		Language:    lang,
	})

	register(&Fmt{
		Name: "black",
		Errorformat: []string{
			`%-GOh no!%.%#`,
			`%-G%\d\+ file%.%#`,
			`%-GAll done!%.%#`,
			`%trror:%f:%l:%c:%m`,
			`%trror:%m`,
			`%m %f`,
			`%-G%.%#`,
		},
		Description: "A uncompromising Python code formatter",
		URL:         "https://github.com/psf/black",
		Language:    lang,
	})

	register(&Fmt{
		Name: "mypy",
		Errorformat: []string{
			`%f:%l:%c: %trror: %m`,
			`%f:%l:%c: %tarning: %m`,
			`%f:%l:%c: %tote: %m`,
			`%f:%l: %trror: %m`,
			`%f:%l: %tarning: %m`,
			`%f:%l: %tote: %m`,
			`%-GFound %\d\+ error%.%#`,
			`%-GSuccess: %.%#`,
		},
		Description: "(mypy --show-column-numbers) Optional static typing for Python",
		URL:         "https://github.com/python/mypy",
		Language:    lang,
	})
}
