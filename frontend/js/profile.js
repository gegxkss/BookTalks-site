// Проверка авторизации и загрузка профиля пользователя
fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
    .then(res => res.json())
    .then(data => {
        if (!data.authenticated) {
            // Перенаправление на /BookTalks-site/signIn.html (корень, а не frontend)
            window.location.href = '/BookTalks-site/signIn.html';
            return;
        }
        // Если авторизован, загружаем данные профиля
        fetch('/BookTalks-site/backend/user_profile_data.php', { credentials: 'same-origin' })
            .then(res => res.json())
            .then(profile => {
                if (profile.error) {
                    // Если ошибка авторизации или другая ошибка backend
                    console.error('Ошибка профиля:', profile.error);
                    // Можно вывести сообщение пользователю
                    return;
                }
                if (profile.user) {
                    document.getElementById('profile-nickname').textContent = profile.user.nickname || '';
                    document.getElementById('profile-email').textContent = profile.user.email || '';
                    document.getElementById('profile-name').textContent = profile.user.first_name || '';
                    document.getElementById('profile-lastname').textContent = profile.user.last_name || '';
                    document.getElementById('profile-sex').textContent = (profile.user.sex === 'male') ? 'мужской' : (profile.user.sex === 'female' ? 'женский' : '');
                    document.getElementById('profile-birthdate').textContent = profile.user.birth_date ? new Date(profile.user.birth_date).toLocaleDateString('ru-RU') : '';
                    if (profile.user.profile_image) {
                        // Всегда абсолютный путь от корня сайта
                        document.getElementById('profile-image').src = '/BookTalks-site/' + profile.user.profile_image.replace(/^\/+/,'')
                    } else {
                        document.getElementById('profile-image').src = '/BookTalks-site/styles/images/icon.png';
                    }
                    // Добавляем вывод значения profile_image в консоль для отладки
                    console.log('Profile image path:', profile.user.profile_image);
                }
                // Динамический вывод книг, рецензий, цитат
                // Книги
                const myBookContainer = document.querySelector('.myBook');
                if (myBookContainer && profile.books) {
                    myBookContainer.innerHTML = '';
                    profile.books.forEach(book => {
                        let coverPath = book.coverimage_filename && book.coverimage_filename !== ''
                            ? '/BookTalks-site/uploads/' + book.coverimage_filename
                            : '/BookTalks-site/uploads/no_cover.png';
                        myBookContainer.innerHTML += `
                            <div class="book">
                                <img src="${coverPath}" alt="Обложка книги" class="book_about" onerror="this.onerror=null;this.src='/BookTalks-site/uploads/no_cover.png';">
                                <p class="name">${book.name}</p>
                            </div>
                        `;
                    });
                }
                // Рецензии
                const reviewsBlock = document.querySelector('.book-resenzii');
                if (reviewsBlock && profile.reviews) {
                    reviewsBlock.innerHTML = '';
                    profile.reviews.forEach(r => {
                        let coverPath = r.coverimage_filename && r.coverimage_filename !== ''
                            ? '/BookTalks-site/uploads/' + r.coverimage_filename
                            : '/BookTalks-site/uploads/no_cover.png';
                        reviewsBlock.innerHTML += `
                            <div class="book-quote">
                                <img src="${coverPath}" class="knizka" alt="Обложка книги" onerror="this.onerror=null;this.src='/BookTalks-site/uploads/no_cover.png';">
                                <div class="book-info">
                                    <div class="name">
                                        <p class="nameBook">${r.book_name}</p>
                                    </div>
                                    <hr class="res-line">
                                    <div class="res-disc">${r.text}</div>
                                </div>
                            </div>
                        `;
                    });
                }
                // Цитаты
                const quotesBlock = document.querySelector('.book-quote');
                if (quotesBlock && profile.quotes) {
                    quotesBlock.innerHTML = '';
                    profile.quotes.forEach(q => {
                        let coverPath = q.coverimage_filename && q.coverimage_filename !== ''
                            ? '/BookTalks-site/uploads/' + q.coverimage_filename
                            : '/BookTalks-site/uploads/no_cover.png';
                        quotesBlock.innerHTML += `
                            <div class="book-quote">
                                <img src="${coverPath}" class="knizka" alt="Обложка книги" onerror="this.onerror=null;this.src='/BookTalks-site/uploads/no_cover.png';">
                                <div class="book-info">
                                    <div class="name">
                                        <p class="nameBook">${q.book_name}</p>
                                    </div>
                                    <hr class="res-line">
                                    <div class="res-disc">${q.text}</div>
                                </div>
                            </div>
                        `;
                    });
                }
                // Можно дополнительно обработать ошибки, если что-то пошло не так
            }).catch(err => {
                console.error('Ошибка при разборе профиля:', err);
                // Можно вывести сообщение пользователю
            });
    });
