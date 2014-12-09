// Return the list of people in the company directory as json.
func PeopleHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Person").
		Order("Name")

	var people []Person
	_, err := q.GetAll(c, &people)
	if err != nil {
		c.Errorf("Error retrieving from datastore %v", err)
		return
	}

	results, err := json.Marshal(people)
	if err != nil {
		c.Errorf("Error while marshalling people %v", err)
		return
	}
	fmt.Fprintf(w, "%s", results)
}
