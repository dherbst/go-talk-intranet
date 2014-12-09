// Add a single person so that we can then use the data viewer to add
func FakePersonAddHandler(w http.ResponseWriter, r *http.Request) {
	c:= appengine.NewContext(r)
	p := Person{
		Name:         "Darrel Herbst",
		Email:        "dherbst@dramafever.com",
		Phone:        "111-555-1212",
		Manager:      "Seung Bak",
		Department:   "Technology",
		Office:       "Narberth",
		Title:        "CTO",
		ThumbnailUrl: "https://plus.google.com/_/focus/photos/public/AIbEiAI...",
	}
	key := datastore.NewKey(c, "Person", p.Email, 0, nil)
	storedPerson, _ := datastore.Put(c, key, &p)
	fmt.Fprintf(w, "Done added person=%v", storedPerson)
}
