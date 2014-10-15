
type Greeting struct {
        Author  string
        Content string
        Date    time.Time
}

func x(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    g := Greeting{
        Content: r.FormValue("content"),
        Date:    time.Now(),
    }	

    rootKey := datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
    key := datastore.NewIncompleteKey(c, "Greeting", rootKey)
	_, err := datastore.Put(c, key, &g)	
}
