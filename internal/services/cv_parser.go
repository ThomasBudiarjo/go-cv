package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type CVParserService struct{}

func NewCVParserService() *CVParserService {
	return &CVParserService{}
}

func (cvp *CVParserService) ParseCV(file multipart.File, header *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(header.Filename))
	
	switch ext {
	case ".pdf":
		return cvp.parsePDF(file)
	case ".docx":
		return cvp.parseDOCX(file)
	case ".txt":
		return cvp.parseText(file)
	default:
		return "", fmt.Errorf("unsupported file format: %s", ext)
	}
}

func (cvp *CVParserService) parsePDF(file multipart.File) (string, error) {
	// TODO: Implement PDF parsing using a library like unidoc/unipdf
	// For now, return a placeholder
	return "PDF parsing not yet implemented. Please use .txt or .docx files for now.", nil
}

func (cvp *CVParserService) parseDOCX(file multipart.File) (string, error) {
	// TODO: Implement DOCX parsing using a library like nguyenthenguyen/docx
	// For now, return a placeholder
	return "DOCX parsing not yet implemented. Please use .txt files for now.", nil
}

func (cvp *CVParserService) parseText(file multipart.File) (string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read text file: %w", err)
	}
	
	return string(content), nil
}