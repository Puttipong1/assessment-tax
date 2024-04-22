package model

type Deduction struct {
	Personal float64
	Donation float64
}

func (d *Deduction) SetDonation(donation float64) {
	d.Donation = donation
}
