package apimanagement_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/apimanagement/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type ApiManagementSubscriptionResource struct {
}

func TestAccApiManagementSubscription_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allow_tracing").HasValue("true"),
				check.That(data.ResourceName).Key("subscription_id").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiManagementSubscription_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("subscription_id").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccApiManagementSubscription_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.update(data, "submitted", true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("submitted"),
				check.That(data.ResourceName).Key("allow_tracing").HasValue("true"),
				check.That(data.ResourceName).Key("subscription_id").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
			),
		},
		{
			Config: r.update(data, "active", true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("active"),
			),
		},
		{
			Config: r.update(data, "suspended", true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("suspended"),
			),
		},
		{
			Config: r.update(data, "cancelled", true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("cancelled"),
			),
		},
		{
			Config: r.update(data, "active", false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allow_tracing").HasValue("false"),
			),
		},
	})
}

func TestAccApiManagementSubscription_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("active"),
				check.That(data.ResourceName).Key("allow_tracing").HasValue("false"),
				check.That(data.ResourceName).Key("subscription_id").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiManagementSubscription_withoutUser(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withoutUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("state").HasValue("active"),
				check.That(data.ResourceName).Key("allow_tracing").HasValue("false"),
				check.That(data.ResourceName).Key("subscription_id").Exists(),
				check.That(data.ResourceName).Key("primary_key").Exists(),
				check.That(data.ResourceName).Key("secondary_key").Exists(),
				check.That(data.ResourceName).Key("user_id").HasValue(""),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiManagementSubscription_withApiId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withApiId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api_id").Exists(),
				check.That(data.ResourceName).Key("product_id").HasValue(""),
				check.That(data.ResourceName).Key("user_id").HasValue(""),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApiManagementSubscription_allApis(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_api_management_subscription", "test")
	r := ApiManagementSubscriptionResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.allApis(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("product_id").HasValue(""),
				check.That(data.ResourceName).Key("user_id").HasValue(""),
			),
		},
		data.ImportStep(),
	})
}

func (ApiManagementSubscriptionResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.SubscriptionID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ApiManagement.SubscriptionsClient.Get(ctx, id.ResourceGroup, id.ServiceName, id.Name)
	if err != nil {
		return nil, fmt.Errorf("reading ApiManagement Subscription (%s): %+v", id, err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (ApiManagementSubscriptionResource) basic(data acceptance.TestData) string {
	template := ApiManagementSubscriptionResource{}.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  user_id             = azurerm_api_management_user.test.id
  product_id          = azurerm_api_management_product.test.id
  display_name        = "Butter Parser API Enterprise Edition"
}
`, template)
}

func (r ApiManagementSubscriptionResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_subscription" "import" {
  subscription_id     = azurerm_api_management_subscription.test.subscription_id
  resource_group_name = azurerm_api_management_subscription.test.resource_group_name
  api_management_name = azurerm_api_management_subscription.test.api_management_name
  user_id             = azurerm_api_management_subscription.test.user_id
  product_id          = azurerm_api_management_subscription.test.product_id
  display_name        = azurerm_api_management_subscription.test.display_name
}
`, r.basic(data))
}

func (r ApiManagementSubscriptionResource) update(data acceptance.TestData, state string, allow_tracing bool) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  user_id             = azurerm_api_management_user.test.id
  product_id          = azurerm_api_management_product.test.id
  display_name        = "Butter Parser API Enterprise Edition"
  state               = "%s"
  allow_tracing       = %t
}
`, r.template(data), state, allow_tracing)
}

func (r ApiManagementSubscriptionResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  user_id             = azurerm_api_management_user.test.id
  product_id          = azurerm_api_management_product.test.id
  display_name        = "Butter Parser API Enterprise Edition"
  state               = "active"
  allow_tracing       = false
}
`, r.template(data))
}

func (r ApiManagementSubscriptionResource) withoutUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  product_id          = azurerm_api_management_product.test.id
  display_name        = "Butter Parser API Enterprise Edition"
  state               = "active"
  allow_tracing       = false
}
`, r.template(data))
}

func (ApiManagementSubscriptionResource) withApiId(data acceptance.TestData) string {
	template := ApiManagementSubscriptionResource{}.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_api" "test" {
  name                = "TestApi"
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  revision            = "1"
  protocols           = ["https"]
  display_name        = "Test API"
  path                = "test"
}

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  api_id              = azurerm_api_management_api.test.id
  display_name        = "Butter Parser API Enterprise Edition"
}
`, template)
}

func (ApiManagementSubscriptionResource) allApis(data acceptance.TestData) string {
	template := ApiManagementSubscriptionResource{}.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_api_management_api" "test" {
  name                = "TestApi"
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  revision            = "1"
  protocols           = ["https"]
  display_name        = "Test API"
  path                = "test"
}

resource "azurerm_api_management_subscription" "test" {
  resource_group_name = azurerm_api_management.test.resource_group_name
  api_management_name = azurerm_api_management.test.name
  display_name        = "Butter Parser API Enterprise Edition"
}
`, template)
}

func (ApiManagementSubscriptionResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_api_management" "test" {
  name                = "acctestAM-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  publisher_name      = "pub1"
  publisher_email     = "pub1@email.com"

  sku_name = "Developer_1"
}

resource "azurerm_api_management_product" "test" {
  product_id            = "test-product"
  api_management_name   = azurerm_api_management.test.name
  resource_group_name   = azurerm_resource_group.test.name
  display_name          = "Test Product"
  subscription_required = true
  approval_required     = false
  published             = true
}

resource "azurerm_api_management_user" "test" {
  user_id             = "acctestuser%d"
  api_management_name = azurerm_api_management.test.name
  resource_group_name = azurerm_resource_group.test.name
  first_name          = "Acceptance"
  last_name           = "Test"
  email               = "azure-acctest%d@example.com"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
