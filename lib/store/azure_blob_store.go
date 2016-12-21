// +build AZURE

package store

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	azure "github.com/azure/azure-sdk-for-go/storage"
	"github.com/cenkalti/backoff"
)