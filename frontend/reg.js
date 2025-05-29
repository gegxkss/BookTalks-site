document.addEventListener('DOMContentLoaded', function() {
    // Добавляем обработчик события на изменение файла
    document.getElementById('profile-upload').addEventListener('change', function(event) {
        const file = event.target.files[0]; // Получаем выбранный файл

        if (file) {
            const reader = new FileReader(); // Создаем FileReader

            reader.onload = function(e) {
                // Когда FileReader закончит чтение файла, обновляем атрибут src изображения
                document.getElementsByClassName('profile-image')[0].src = e.target.result;
            }

            reader.readAsDataURL(file); // Читаем файл как Data URL
        }
    });
});


function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
}

function register() {
    let isValid = true; // Объявляем isValid здесь

    // 1. Получаем значения полей
    const firstName = document.getElementById('firstName').value.trim();
    const lastName = document.getElementById('lastName').value.trim();
    const sex = document.getElementById('sex').value;
    const birthDate = document.getElementById('birthDate').value;
    const nickname = document.getElementById('nickname').value.trim();
    const email = document.getElementById('email').value.trim();
    const password = document.getElementById('password').value;

    // 2. Валидация данных (клиентская)

    // Сбрасываем предыдущие ошибки
    document.querySelectorAll('.error').forEach(el => el.textContent = '');

    if (firstName === "") {
        document.getElementById('firstNameError').textContent = 'Пожалуйста, введите имя.';
        isValid = false;
    }

    if (lastName === "") {
        document.getElementById('lastNameError').textContent = 'Пожалуйста, введите фамилию.';
        isValid = false;
    }

    if (sex === "") {
        document.getElementById('sexError').textContent = 'Пожалуйста, выберите пол.';
        isValid = false;
    }

    if (birthDate === "") {
        document.getElementById('birthDateError').textContent = 'Пожалуйста, введите дату рождения.';
        isValid = false;
    }

    if (nickname === "") {
        document.getElementById('nicknameError').textContent = 'Пожалуйста, введите никнейм.';
        isValid = false;
    }

    if (email === "") {
        document.getElementById('emailError').textContent = 'Пожалуйста, введите почту.';
        isValid = false;
    } else if (!validateEmail(email)) {
        document.getElementById('emailError').textContent = 'Пожалуйста, введите корректную почту.';
        isValid = false;
    }

    if (password === "") {
        document.getElementById('passwordError').textContent = 'Пожалуйста, введите пароль.';
        isValid = false;
    } else if (password.length < 6) {
        document.getElementById('passwordError').textContent = 'Пароль должен быть не менее 6 символов.';
        isValid = false;
    }

    console.log('Функция register() вызвана.');

    if (!isValid) {
        return;
    }

    // Создаем объект FormData для отправки данных формы
    const formData = new FormData();

    // Добавляем данные из полей формы
    formData.append('nickname', nickname);
    formData.append('first_name', firstName);
    formData.append('last_name', lastName);
    formData.append('sex', sex);
    formData.append('birth_date', birthDate);
    formData.append('email', email);
    formData.append('password', password);

    // Получаем файл изображения
    const profileImageFile = document.getElementById('profile-upload').files[0];
    if (profileImageFile) {
        formData.append('profile_image', profileImageFile); // Добавляем изображение в FormData
    }
    console.log('Отправляю данные на сервер...');

    function validateDate(dateString) {
  // Пример валидации для формата "YYYY-MM-DD"
  const regex = /^\d{4}-\d{2}-\d{2}$/;
  if (!regex.test(dateString)) {
    return false; // Неверный формат
  }

  // Дополнительная проверка на корректность даты (например, на существование 31 февраля)
  const date = new Date(dateString);
  if (isNaN(date.getTime())) {
    return false; // Некорректная дата
  }

  return true; // Верный формат
}

const birthDateInput = document.getElementById('birthDate');
birthDateInput.addEventListener('blur', () => {
  const birthDate = birthDateInput.value;
  if (!validateDate(birthDate)) {
    alert('Неверный формат даты рождения. Используйте формат YYYY-MM-DD.');
    birthDateInput.value = ''; // Очищаем поле
  }
});

    // Отправляем данные на сервер (AJAX)
    fetch('/BookTalks-site/backend/register.php', {
            method: 'POST',
            body: formData // Отправляем FormData
        })
    .then(response => {
        if (response.ok) {
            // Успешная регистрация
            response.json().then(data => {
                if (data && data.profile_image_url) {
                    // Сохраняем URL изображения профиля в localStorage
                    localStorage.setItem('profileImageUrl', data.profile_image_url);
                }
                alert('Регистрация прошла успешно!');
                // Перенаправляем пользователя на главную страницу
                window.location.href = '/'; //  Главная страница
            });
        } else {
            // Обработка ошибок от сервера
            response.text().then(text => {
                console.error('Ошибка регистрации:', text);
                alert('Ошибка регистрации: ' + text);
            });
        }
    })
    .catch(error => {
        console.error('Ошибка:', error);
        alert('Произошла ошибка при отправке запроса.');
    });
}