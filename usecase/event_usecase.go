package usecase

import (
	"desafio/domain"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type EventUseCases struct {
	data domain.Data
}

func NewEventUseCases() *EventUseCases {
	return &EventUseCases{}
}

func (uc *EventUseCases) LoadData(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &uc.data); err != nil {
		return err
	}

	return nil
}

func (uc *EventUseCases) GetEvents() []domain.Event {
	return uc.data.Events
}

func (uc *EventUseCases) GetEventByID(id int) (domain.Event, error) {
	for _, event := range uc.data.Events {
		if event.ID == id {
			return event, nil
		}
	}
	return domain.Event{}, errors.New("event not found")
}

func (uc *EventUseCases) GetSpotsByEventID(eventID int) ([]domain.Spot, error) {
	var spots []domain.Spot

	log.Printf("No spots found for event ID: %d", eventID)
	for _, spot := range uc.data.Spots {
		if spot.EventID == eventID {
			spots = append(spots, spot)
		}
	}
	if len(spots) == 0 {
		return nil, errors.New("no spots found for this event")
	}
	return spots, nil
}

func (uc *EventUseCases) ReserveSpot(eventID int, spots []string) ([]domain.Spot, error) {
	var reservedSpots []domain.Spot

	for _, name := range spots {
		found := false
		for i, spot := range uc.data.Spots {
			if spot.EventID == eventID && spot.Name == name {
				if spot.Status == domain.SpotStatusReserved {
					return nil, fmt.Errorf("spot %s is already reserved", name)
				}
				uc.data.Spots[i].Status = domain.SpotStatusReserved
				reservedSpots = append(reservedSpots, spot)
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("spot %s not found for event ID %d", name, eventID)
		}
	}

	return reservedSpots, nil
}
