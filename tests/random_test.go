package tests

import (
	"gorm.io/gorm"
	"pasteProject/errs"
	"pasteProject/models"
	"pasteProject/repositories"
	"reflect"
	"testing"
)

// 测试启动函数
func TestMain(m *testing.M) {
	m.Run()
}

// 测试用户工作流
func TestUserWorkFlow(t *testing.T) {
	t.Run("create_user", testUserRepo_CreateU)
	t.Run("select_user", testUserRepo_SelectU)
	t.Run("update_user", testUserRepo_UpdateU)
	t.Run("delete_user", testUserRepo_DeleteU)
}

// 测试粘帖工作流
func TestPasteWorkFlow(t *testing.T) {
	t.Run("create_paste", testPasteRepo_CreateP)
	t.Run("select_pastes", testPasteRepo_SelectPS)
	t.Run("delete_paste", testPasteRepo_DeleteP)
	t.Run("delete_pastes", testPasteRepo_DeletePS)
}

// 测试创建用户
func testUserRepo_CreateU(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"testU1", fields{errs.GetDB()}, args{&models.User{UserName: "name1", EncryptedPwd: "1"}}, false},
		{"testU2", fields{errs.GetDB()}, args{&models.User{UserName: "name2", EncryptedPwd: "12"}}, false},
		{"testU3", fields{errs.GetDB()}, args{&models.User{UserName: "name3", EncryptedPwd: "123"}}, false},
		{"testU4", fields{errs.GetDB()}, args{&models.User{UserName: "name4", EncryptedPwd: "1234"}}, false},
		{"testU5", fields{errs.GetDB()}, args{&models.User{UserName: "name5", EncryptedPwd: "12345"}}, false},
		{"testU6", fields{errs.GetDB()}, args{&models.User{UserName: "name6", EncryptedPwd: "123456"}}, false},
		{"testU7", fields{errs.GetDB()}, args{&models.User{UserName: "name7", EncryptedPwd: "1234567"}}, false},
		{"testU8", fields{errs.GetDB()}, args{&models.User{UserName: "name8", EncryptedPwd: "12345678"}}, false},
		{"testU9", fields{errs.GetDB()}, args{&models.User{UserName: "name9", EncryptedPwd: "123456789"}}, false},
		{"testU10", fields{errs.GetDB()}, args{&models.User{UserName: "name10", EncryptedPwd: "12345678910"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &repositories.UserRepo{
				Db: tt.fields.Db,
			}
			if err := u.CreateU(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateU() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试删除用户
func testUserRepo_DeleteU(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"testU6", fields{errs.GetDB()}, args{&models.User{UserName: "name6", EncryptedPwd: "123456"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &repositories.UserRepo{
				Db: tt.fields.Db,
			}
			if err := u.DeleteU(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("DeleteU() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试查询用户
func testUserRepo_SelectU(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{"testU8", fields{errs.GetDB()}, args{"name8"}, &models.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &repositories.UserRepo{
				Db: tt.fields.Db,
			}
			got, err := u.SelectU(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.UserName != "name8" {
				t.Errorf("SelectU() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试更新用户信息
func testUserRepo_UpdateU(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"testU4", fields{errs.GetDB()}, args{&models.User{UserName: "name4", EncryptedPwd: "123456789"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &repositories.UserRepo{
				Db: tt.fields.Db,
			}
			if err := u.UpdateU(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateU() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试创建粘帖
func testPasteRepo_CreateP(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		paste *models.Paste
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"testP1", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga1", Poster: "name1"}}, false}, {"testP2", fields{errs.GetDB()}, args{&models.Paste{Things: "agmmgyumyu2", Poster: "name2"}}, false}, {"testP3", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga3", Poster: "name3"}}, false}, {"testP4", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga4", Poster: "name4"}}, false}, {"testP5", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga5", Poster: "name5"}}, false}, {"testP6", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga6", Poster: "name6"}}, false}, {"testP7", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga7", Poster: "name7"}}, false}, {"testP8", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga8", Poster: "name8"}}, false}, {"testP9", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga9", Poster: "name9"}}, false}, {"testP10", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga10", Poster: "name10"}}, false}, {"testP11", fields{errs.GetDB()}, args{&models.Paste{Things: "asgargarga10", Poster: "name10"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repositories.PasteRepo{
				Db: tt.fields.Db,
			}
			if err := p.CreateP(tt.args.paste); (err != nil) != tt.wantErr {
				t.Errorf("CreateP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试删除粘帖
func testPasteRepo_DeleteP(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		posterName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"testP1", fields{errs.GetDB()}, args{"name1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repositories.PasteRepo{
				Db: tt.fields.Db,
			}
			if err := p.DeleteP(tt.args.posterName); (err != nil) != tt.wantErr {
				t.Errorf("DeleteP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试删除所有粘帖
func testPasteRepo_DeletePS(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"testP10", fields{errs.GetDB()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repositories.PasteRepo{
				Db: tt.fields.Db,
			}
			if err := p.DeletePS(); (err != nil) != tt.wantErr {
				t.Errorf("DeletePS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试获取所有粘帖
func testPasteRepo_SelectPS(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.Paste
	}{
		{"testP10", fields{errs.GetDB()}, []models.Paste{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &repositories.PasteRepo{
				Db: tt.fields.Db,
			}
			if got := p.SelectPS(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectPS() = %v, want %v", got, tt.want)
			}
		})
	}
}
