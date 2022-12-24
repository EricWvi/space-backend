package service

import (
	"github.com/space-backend/model"
	"testing"
)

func TestUploadFile(t *testing.T) {
	type args struct {
		fileType int
		name     string
		content  []byte
	}
	tests := []struct {
		name     string
		args     args
		wantLink string
		wantErr  bool
	}{
		{
			name: "1",
			args: args{
				fileType: model.LocalOSS,
				name:     "1234test",
				content:  []byte("content"),
			},
			wantLink: "12/34/1234test",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLink, err := UploadFile(tt.args.fileType, tt.args.name, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLink != tt.wantLink {
				t.Errorf("UploadFile() gotLink = %v, want %v", gotLink, tt.wantLink)
			}
			content, err := DownloadFile(tt.args.fileType, tt.wantLink)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(content) != string(tt.args.content) {
				t.Errorf("DownloadFile() gotContent = %v, want %v", content, string(content))
			}
		})
	}
}
