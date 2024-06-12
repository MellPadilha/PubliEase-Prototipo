package cache

import (
	"os"
)

type Cache struct {
	Name    string
	Path    string
	Content []byte
}

var path = "./cache/"

func New() *Cache {
	return &Cache{}
}

// CreateCacheDir create the cache directory
func (c *Cache) CreateCacheDir() error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// CreateCacheFile create a cache file
func (c *Cache) CreateCacheFile() error {
	c.Path = path + c.Name

	err := os.WriteFile(path+c.Name, c.Content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetCacheFile get cache file from cache directory
func (c *Cache) GetCacheFile() error {
	var err error
	c.Path = path + c.Name

	c.Content, err = os.ReadFile(c.Path)
	if err != nil {
		return err
	}

	return nil
}

// VerifyFile verify if file exists
func (c *Cache) VerifyFile() bool {
	c.Path = path + c.Name
	_, err := os.Stat(c.Path)
	if err != nil {
		return false
	}

	return true
}
