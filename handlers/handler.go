package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aliffatulmf/pink/requests"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

const expiryDuration = 720 * time.Hour  // 30 days
const cleanupInterval = 5 * time.Minute // 5 minutes

var domainCache = cache.New(expiryDuration, cleanupInterval)

func VerifyDomain(c *gin.Context) {
	domainName, exists := c.GetQuery("name")
	if !exists {
		sendError(c, http.StatusBadRequest, "query name is required.")
		return
	}

	var parsedDomain *requests.Domain
	cachedDomain, found := domainCache.Get(domainName)
	if found {
		c.JSON(http.StatusOK, cachedDomain)
		return
	} else {
		domainResponse, err := requests.NewRequest(domainName)
		if err != nil {
			sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to check the names: %v", err))
			return
		}

		parsedDomain = requests.ParseRecord(domainResponse)
		domainCache.Set(domainName, parsedDomain, cache.DefaultExpiration)
	}

	c.JSON(http.StatusOK, parsedDomain)
}

func sendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

func ListDomains(c *gin.Context) {
	var list []interface{}

	if domainCache.ItemCount() == 0 {
		c.String(http.StatusOK, "no domain cached")
		return
	}

	for _, domain := range domainCache.Items() {
		obj := domain.Object.(*requests.Domain)
		list = append(list, obj.Domain)
	}

	c.JSON(http.StatusOK, gin.H{
		"domain": list,
	})
}
