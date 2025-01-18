class User < ApplicationRecord
  has_secure_password

  enum :role, { student: 0, tutor: 1 }

  validates :user_name, presence: true
  validates :email, presence: true, uniqueness: true
  validates :password, presence: true
end
