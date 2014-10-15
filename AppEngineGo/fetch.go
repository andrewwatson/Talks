

rootKey := datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)

q := datastore.NewQuery("Greeting").Ancestor(rootKey).Order("-Date").Limit(10)

greetings := make([]Greeting, 0, 10)

if _, err := q.GetAll(c, &greetings); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}