package storage

import (
	"context"
	"crypto/rand"
	"log/slog"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"golang.org/x/crypto/bcrypt"

	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/internal/errs"
)

type Storage struct {
	db  *pg.DB
	log *slog.Logger
}

func NewStorage(cfg *config.Config, log *slog.Logger) (*Storage, error) {
	storage := &Storage{
		log: log,
	}
	opt, err := pg.ParseURL(cfg.DBConnectString)
	if err != nil {
		return nil, err
	}
	storage.log.Info("Opening DB")
	storage.db = pg.Connect(opt)
	err = storage.createSchema()
	if err != nil {
		return nil, err
	}
	return storage, nil
}

func (d *Storage) createSchema() error {
	models := []interface{}{
		(*User)(nil),
		(*Secret)(nil),
	}

	for _, model := range models {
		err := d.db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Storage) Close() {
	d.log.Info("Closing DB")
	d.db.Close()
}

func (d *Storage) AddUser(ctx context.Context, user *User) (int64, error) {
	salt, err := generateSalt()
	if err != nil {
		return -1, err
	}
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	user.Salt = salt
	user.HashedPassword = hashedPassword
	_, err = d.db.ModelContext(ctx, user).Insert()
	if pgErr, ok := err.(pg.Error); ok {
		if pgErr.IntegrityViolation() {
			return -1, errs.ErrConflict
		} else {
			return -1, err
		}
	}
	return user.ID, nil
}

// Проверка на совпадение логина и пароля. Возвращает userID в случае совпадения и -1 в противном случае.
func (d *Storage) CheckLoginPassword(ctx context.Context, user *User) (int64, error) {
	// Проверка, что логин присутствует в базе ...
	err := d.db.ModelContext(ctx, user).Where("name = ?", user.Name).Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return -1, nil
		} else {
			return 0, err
		}
	}
	// ... логин есть. Теперь проверим совпадение паролей.
	salted := append([]byte(user.Password), user.Salt...)
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return -1, nil
	}
	return user.ID, nil
}

func (d *Storage) SaveSecret(ctx context.Context, secret Secret) error {
	tx, err := d.db.BeginContext(ctx)
	if err != nil {
		return err
	}
	defer tx.Close()

	_, err = tx.ModelContext(ctx, &secret).Insert()
	if err != nil {
		if pgErr, ok := err.(pg.Error); ok {
			if pgErr.IntegrityViolation() {
				return errs.ErrConflict
			}
		} else {
			return err
		}
	}

	return tx.Commit()
}

func (d *Storage) GetSecret(ctx context.Context, userID int64, name string) (*Secret, error) {
	secret := Secret{}
	err := d.db.ModelContext(ctx, (*Secret)(nil)).Where("user_id = ?", userID).Where("secret_name = ?", name).Select(&secret)
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
