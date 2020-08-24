package goshopify

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func RiskTests(t *testing.T, risk Risk) {
	// Check that the ID is assigned to the returned risk
	expectedID := int64(284138680)
	if risk.ID != expectedID {
		t.Errorf("Risk.ID returned %+v, expected %+v", risk.ID, expectedID)
	}

	// Check that the OrderID value is assigned to the returned risk
	expectedOrderID := int64(450789469)
	if risk.OrderID != expectedOrderID {
		t.Errorf("Risk.OrderID returned %+v, expected %+v", risk.OrderID, expectedOrderID)
	}

	// Check that the External value is assigned to the returned risk
	expectedSource := "External"
	if risk.Source != expectedSource {
		t.Errorf("risk.Source returned %+v, expected %+v", risk.Source, expectedSource)
	}

	// Check that the Score value is assigned to the returned risk
	expectedScore := "1.0"
	if risk.Score != expectedScore {
		t.Errorf("risk.Score returned %+v, expected %+v", risk.Score, expectedScore)
	}

	// Check that the Recommendation value is assigned to the returned risk
	expectedRecommendation := "cancel"
	if risk.Recommendation != expectedRecommendation {
		t.Errorf("risk.Recommendation returned %+v, expected %+v", risk.Recommendation, expectedRecommendation)
	}

	// Check that the Display value is assigned to the returned risk
	expectedDisplay := true
	if risk.Display != expectedDisplay {
		t.Errorf("risk.Display returned %+v, expected %+v", risk.Display, expectedDisplay)
	}

	// Check that the Message value is assigned to the returned risk
	expectedMessage := "This order was placed from a proxy IP"
	if risk.Message != expectedMessage {
		t.Errorf("risk.Message returned %+v, expected %+v", risk.Message, expectedMessage)
	}

	// Check that the MerchantMessage value is assigned to the returned risk
	expectedMerchantMessage := "This order was placed from a proxy IP"
	if risk.MerchantMessage != expectedMerchantMessage {
		t.Errorf("risk.MerchantMessage returned %+v, expected %+v", risk.MerchantMessage, expectedMerchantMessage)
	}
}

func TestRiskList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/1/risks.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("risks.json")))

	risks, err := client.Risk.List(1, nil)
	if err != nil {
		t.Fatalf("Risk.List returned error: %v", err)
	}

	for _, risk := range risks {
		RiskTests(t, risk)
	}
}

func TestRiskGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/1/risks/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("risk.json")))

	risk, err := client.Risk.Get(1, 1, nil)
	if err != nil {
		t.Fatalf("Risk.Get returned error: %v", err)
	}

	RiskTests(t, *risk)
}

func TestRiskCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/1/risks.json", client.pathPrefix),
		httpmock.NewBytesResponder(201, loadFixture("risk.json")))

	risk := Risk{
		Message:        "This order came from an anonymous proxy",
		Recommendation: "cancel",
		Score:          "1.0",
		Source:         "External",
		CauseCancel:    true,
		Display:        true,
	}

	result, err := client.Risk.Create(1, risk)
	if err != nil {
		t.Fatalf("Risk.Create returned error: %+v", err)
	}
	RiskTests(t, *result)
}

func TestRiskUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/1/risks/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("risk.json")))

	risk := Risk{
		ID:             1,
		Message:        "This order came from an anonymous proxy",
		Recommendation: "cancel",
		Score:          "1.0",
		Source:         "External",
		CauseCancel:    true,
		Display:        true,
	}

	result, err := client.Risk.Update(1, risk)
	if err != nil {
		t.Fatalf("Risk.Update returned error: %+v", err)
	}
	RiskTests(t, *result)
}

func TestRiskDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/1/risks/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("risk.json")))

	err := client.Risk.Delete(1, 1)
	if err != nil {
		t.Fatalf("Risk.Delete returned error: %+v", err)
	}
}
