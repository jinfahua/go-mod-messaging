/********************************************************************************
 *  Copyright 2020 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package pkg

import "fmt"

// CertificateErr represents an error associated with interacting with a Certificate.
type CertificateErr struct {
	description string
}

func (ce CertificateErr) Error() string {
	return fmt.Sprintf("Unable to process certificate properties: %s", ce.description)
}

// NewCertificateErr constructs a new CertificateErr
func NewCertificateErr(message string) CertificateErr {
	return CertificateErr{description: message}
}

// BrokerURLErr represents an error associated parsing a broker's URL.
type BrokerURLErr struct {
	description string
}

func (bue BrokerURLErr) Error() string {
	return fmt.Sprintf("Unable to process broker URL: %s", bue.description)
}

// NewBrokerURLErr constructs a new BrokerURLErr
func NewBrokerURLErr(description string) BrokerURLErr {
	return BrokerURLErr{description: description}
}
