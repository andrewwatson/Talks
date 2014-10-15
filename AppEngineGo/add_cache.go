
// Create an Item
item := &memcache.Item{
    Key:   "lyric",
    Value: []byte("Oh, give me a home"),
    Expiration: time.Hour, // HL
}

// Add the item to the memcache, if the key does not already exist
if err := memcache.Add(c, item); err == memcache.ErrNotStored {
    c.Infof("item with key %q already exists", item.Key)
} else if err != nil {
    c.Errorf("error adding item: %v", err)
}
