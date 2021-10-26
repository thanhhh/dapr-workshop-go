package finecollection

type FineCalculator interface {
	CalculateFine(licenseKey string, violationInKmh int) (int, error)
}
