package utils

import (
	"fmt"
	"math"
	"strings"
)

type PaginationRequest struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`
}

type PaginationResponse struct {
	TotalCount int    `json:"total_count"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	HasMore    bool   `json:"has_more"`
	OrderBy    string `json:"order_by"`
	OrderType  string `json:"order_type"`
}

type Default struct {
	Size      int
	OrderBy   string
	OrderType string
}

type PageData struct {
	Paging  *PaginationRequest
	Default Default
}

func (p *PageData) PaginationQueryBuilder(isStandalone bool) string {

	if p.Paging == nil {
		return ""
	}
	if p.Paging.Page <= 0 {
		p.Paging.Page = 1
	}
	if p.Paging.Size <= 0 || p.Default.Size == 0 || p.Default.OrderBy == "" || p.Default.OrderType == "" {
		return ""
	}
	if p.Paging.OrderBy == "" {
		p.Paging.OrderBy = p.Default.OrderBy
	}
	if p.Paging.OrderType == "" {
		p.Paging.OrderType = p.Default.OrderType
	}

	p.Paging.OrderBy = strings.ReplaceAll(p.Paging.OrderBy, " ", "")
	p.Paging.OrderType = strings.ReplaceAll(p.Paging.OrderType, " ", "")

	limit := p.Paging.Size
	offset := (p.Paging.Page - 1) * p.Paging.Size

	return fmt.Sprintf(` ORDER BY %s %s LIMIT %v OFFSET %v`, p.Paging.OrderBy, p.Paging.OrderType, limit, offset)
}

func (p PageData) GetPaginationResponse(totalData int) *PaginationResponse {
	var totalPages int
	if p.Paging.Size > 0 {
		totalPages = int(math.Ceil(float64(totalData) / float64(p.Paging.Size)))
	} else {
		totalPages = 1
	}
	var hasMore bool
	if totalPages-p.Paging.Page < 1 {
		hasMore = false
	} else {
		hasMore = true
	}
	return &PaginationResponse{
		TotalCount: totalData,
		TotalPages: totalPages,
		Page:       p.Paging.Page,
		Size:       p.Paging.Size,
		HasMore:    hasMore,
		OrderBy:    p.Paging.OrderBy,
		OrderType:  p.Paging.OrderType,
	}
}
