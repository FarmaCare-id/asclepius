package dto

type (
	// CreateAppointmentRequest CreateAppointmentRequest
	CreateAppointmentRequest struct {
		HealthcareWorkerID uint `json:"healthcare_worker_id" validate:"required"`
		Complaint string `json:"complaint" validate:"required"`
		ScheduleDate string `json:"schedule_date" validate:"required"`
		ScheduleTime string `json:"schedule_time" validate:"required"`
		Location string `json:"location" validate:"required"`
		Status string `json:"status" validate:"required"`
		Note string `json:"note"`
	}

	// CreateAppointmentResponse CreateAppointmentResponse
	CreateAppointmentResponse struct {
		ID uint `json:"id"`
		PatientID uint `json:"patient_id"`
		HealthcareWorkerID uint `json:"healthcare_worker_id"`
		Complaint string `json:"complaint"`
		ScheduleDate string `json:"schedule_date"`
		ScheduleTime string `json:"schedule_time"`
		Location string `json:"location"`
		Status string `json:"status"`
		Note string `json:"note"`
	}

	// GetAppointmentResponse GetAppointmentResponse
	GetAppointmentResponse struct {
		ID uint `json:"id"`
		PatientID uint `json:"patient_id"`
		HealthcareWorkerID uint `json:"healthcare_worker_id"`
		Complaint string `json:"complaint"`
		ScheduleDate string `json:"schedule_date"`
		ScheduleTime string `json:"schedule_time"`
		Location string `json:"location"`
		Status string `json:"status"`
		Note string `json:"note"`
	}
)