
// Get the item from the memcache
if item, err := memcache.Get(c, "lyric"); err == memcache.ErrCacheMiss {
    c.Infof("item not in the cache")
} else if err != nil {
    c.Errorf("error getting item: %v", err)
} else {
    c.Infof("the lyric is %q", item.Value)
}