package main

import (
	. "github.com/theplant/htmlgo"
)

func homepageHTML() HTMLComponent {
	return HTML(
		Head(
			Meta().Charset("utf-8"),
			Meta().Name("viewport").Content("width=device-width, initial-scale=1"),
			Title(s.RelayName),
			Script("").Src("https://cdn.tailwindcss.com"),
		),
		Body(
			H1("create a group").Class("text-xl mb-2"),
			Form(
				Label("group owner:").For("npub").Class("mr-1 mt-4 block"),
				Input("").Id("npub").Placeholder("npub1...").Name("pubkey").Class("w-96 px-4 py-2 outline-0 bg-stone-100"),
				Label("group name:").For("name").Class("mr-1 mt-4 block"),
				Input("").Id("name").Placeholder("my little group").Name("name").Class("w-96 px-4 py-2 outline-0 bg-stone-100"),

				// Adding the privacy field with Disabled
				Label("privacy:").For("privacy").Class("mr-1 mt-4 block"),
				Select(
					Option("public").Value("public"),
					Option("private").Value("private"),
				).Id("privacy").Name("privacy").Class("w-96 px-4 py-2 outline-0 bg-stone-100").Disabled(true),

				// Adding the type field with Disabled
				Label("type:").For("type").Class("mr-1 mt-4 block"),
				Select(
					Option("open").Value("open"),
					Option("closed").Value("closed"),
				).Id("type").Name("type").Class("w-96 px-4 py-2 outline-0 bg-stone-100").Disabled(true),

				Button("create").Class("block rounded mt-4 px-4 py-2 bg-emerald-500 text-white hover:bg-emerald-300 transition-colors"),
			).Action("/create").Method("POST"),
		).Class("mx-4 my-6"),
	)
}
