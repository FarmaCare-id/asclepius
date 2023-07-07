package appointment

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
	"strconv"
)

type (
	ViewService interface {
		CreateAppointment(ctx dto.SessionContext, req dto.CreateAppointmentRequest) (dto.CreateAppointmentResponse, error)
		GetAllAppointmentByUserID(userId string) ([]dto.GetAppointmentResponse, error)
		GetAllAppointmentForCurrentUser(ctx dto.SessionContext) ([]dto.GetAppointmentResponse, error)
	}

	viewService struct {
		repository repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateAppointment(ctx dto.SessionContext, req dto.CreateAppointmentRequest) (dto.CreateAppointmentResponse, error) {
	var (
		res dto.CreateAppointmentResponse
	)

	appointment := models.Appointment{
		PatientID: ctx.User.ID,
		HealthcareWorkerID: req.HealthcareWorkerID,
		Complaint: req.Complaint,
		ScheduleDate: req.ScheduleDate,
		ScheduleTime: req.ScheduleTime,
		Location: req.Location,
		Status: req.Status,
		Note: req.Note,
	}

	err := v.repository.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return res, err
	}

	res = dto.CreateAppointmentResponse{
		ID: appointment.ID,
		PatientID: appointment.PatientID,
		HealthcareWorkerID: appointment.HealthcareWorkerID,
		Complaint: appointment.Complaint,
		ScheduleDate: appointment.ScheduleDate,
		ScheduleTime: appointment.ScheduleTime,
		Location: appointment.Location,
		Status: appointment.Status,
		Note: appointment.Note,
	}

	return res, nil
}

func (v *viewService) GetAllAppointmentByUserID(userId string) ([]dto.GetAppointmentResponse, error) {
	var (
		res []dto.GetAppointmentResponse
	)

	cid, err := strconv.Atoi(userId)
	if err != nil {
		return res, err
	}

	appointments, err := v.repository.AppointmentRepository.GetAllAppointmentByUserID(uint(cid))
	if err != nil {
		return res, err
	}

	for _, appointment:= range appointments {
		// drug, _ := v.repository.DrugRepository.GetDrugByID(userDrug.DrugID)
		res = append(res, dto.GetAppointmentResponse{
			ID: appointment.ID,
			PatientID: appointment.PatientID,
			HealthcareWorkerID: appointment.HealthcareWorkerID,
			Complaint: appointment.Complaint,
			ScheduleDate: appointment.ScheduleDate,
			ScheduleTime: appointment.ScheduleTime,
			Location: appointment.Location,
			Status: appointment.Status,
			Note: appointment.Note,
		})
	}

	return res, nil
}

func (v *viewService) GetAllAppointmentForCurrentUser(ctx dto.SessionContext) ([]dto.GetAppointmentResponse, error) {
	var (
		res []dto.GetAppointmentResponse
	)

	appointments, err := v.repository.AppointmentRepository.GetAllAppointmentByUserID(ctx.User.ID)
	if err != nil {
		return res, err
	}

	for _, appointment:= range appointments {
		// drug, _ := v.repository.DrugRepository.GetDrugByID(userDrug.DrugID)
		res = append(res, dto.GetAppointmentResponse{
			ID: appointment.ID,
			PatientID: appointment.PatientID,
			HealthcareWorkerID: appointment.HealthcareWorkerID,
			Complaint: appointment.Complaint,
			ScheduleDate: appointment.ScheduleDate,
			ScheduleTime: appointment.ScheduleTime,
			Location: appointment.Location,
			Status: appointment.Status,
			Note: appointment.Note,
		})
	}

	return res, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}