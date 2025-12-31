package service

import (
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
	"time"

	"gorm.io/gorm"
)

type AdminService struct{}

// --- 角色管理 ---

// CreateRole 创建角色
func (s *AdminService) CreateRole(role *model.SysRole) error {
	return global.DB.Create(role).Error
}

// UpdateRole 修改角色
func (s *AdminService) UpdateRole(role *model.SysRole) error {
	return global.DB.Model(&model.SysRole{}).Where("id = ?", role.ID).Updates(role).Error
}

// DeleteRole 删除角色
func (s *AdminService) DeleteRole(id int64) error {
	return global.DB.Delete(&model.SysRole{}, id).Error
}

// ListRoles 角色列表
func (s *AdminService) ListRoles() ([]model.SysRole, error) {
	var roles []model.SysRole
	err := global.DB.Find(&roles).Error
	return roles, err
}

// --- 菜单管理 (权限) ---

// CreateMenu 创建菜单
func (s *AdminService) CreateMenu(menu *model.SysMenu) error {
	return global.DB.Create(menu).Error
}

// ListMenus 获取所有菜单
func (s *AdminService) ListMenus() ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := global.DB.Order("sort asc").Find(&menus).Error
	return menus, err
}

// BindRoleMenu 为角色分配菜单
func (s *AdminService) BindRoleMenu(roleID int64, menuIDs []int64) error {
	tx := global.DB.Begin()
	// 1. 删除旧关联
	if err := tx.Where("role_id = ?", roleID).Delete(&model.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 2. 添加新关联
	var roleMenus []model.SysRoleMenu
	for _, mid := range menuIDs {
		roleMenus = append(roleMenus, model.SysRoleMenu{
			RoleID: roleID,
			MenuID: mid,
		})
	}
	if len(roleMenus) > 0 {
		if err := tx.Create(&roleMenus).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// --- 用户管理 ---

// ListUsers 获取用户列表 (支持搜索)
func (s *AdminService) ListUsers(page, size int, keyword string) ([]model.SysUser, int64, error) {
	var users []model.SysUser
	var total int64
	db := global.DB.Model(&model.SysUser{})

	if keyword != "" {
		db = db.Where("username LIKE ? OR mobile LIKE ? OR real_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	db.Count(&total)
	err := db.Offset((page - 1) * size).Limit(size).Find(&users).Error
	return users, total, err
}

// FreezeUser 冻结/解冻用户
func (s *AdminService) FreezeUser(id int64, status int) error {
	return global.DB.Model(&model.SysUser{}).Where("id = ?", id).Update("status", status).Error
}

// AssignRole 为用户分配角色
func (s *AdminService) AssignRole(userID int64, roleCode string) error {
	return global.DB.Model(&model.SysUser{}).Where("id = ?", userID).Update("role", roleCode).Error
}

// UpdateUserBalance 为用户充值/扣费 (管理员)
func (s *AdminService) UpdateUserBalance(userID int64, amount float64) error {
	if amount == 0 {
		return nil
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 更新余额
		if err := tx.Model(&model.SysUser{}).Where("id = ?", userID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 2. 记录流水
		remark := "管理员调整余额"
		if amount > 0 {
			remark = "系统充值"
		} else {
			remark = "系统扣除"
		}

		transaction := model.SysTransaction{
			UserID:    userID,
			Type:      3, // 3: 系统调整/充值
			Amount:    amount,
			Remark:    remark,
			CreatedAt: time.Now(),
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
}
