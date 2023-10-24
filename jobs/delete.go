package jobs

import (
	"fmt"
)

func (c *cronJob) DeleteRole() {
	c.cron.
		Every(1).
		Day().
		At("12:00").
		Do(func() {
			// service.Role().DeleteMultiple(c.ctx.DB)
			fmt.Println("do stuff")
		})
}
