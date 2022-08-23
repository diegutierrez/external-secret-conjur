/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package register

// packages imported here are registered to the controller schema.
// nolint:revive
import (
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/akeyless"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/alibaba"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/aws"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/azure/keyvault"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/fake"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/gcp/secretmanager"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/gitlab"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/ibm"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/kubernetes"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/onepassword"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/oracle"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/senhasegura"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/vault"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/webhook"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/yandex/certificatemanager"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/yandex/lockbox"
	_ "github.com/diegutierrez/external-secret-conjur/pkg/provider/conjur"
)
