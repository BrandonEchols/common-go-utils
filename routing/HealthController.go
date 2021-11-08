/**
 * Copyright 2017 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */

package routing

import (
	"net/http"
)

/*
	This class is used for hosting a 'health' point that is useful for determining if the service is up and routing
	is working as expected.
*/
type IHealthController interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

//Implements IHealthController
type healthController struct{}

/*
	Returns an implementation of IHealthController
*/
func GetHealthController() IHealthController {
	return &healthController{}
}

/*
	This API endpoint simply returns a UTF-8 character response: "Alive".
*/
func (this *healthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	result := []byte("Alive")

	w.Header().Set("Content-Type", "charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
