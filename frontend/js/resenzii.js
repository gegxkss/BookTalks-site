document.addEventListener('DOMContentLoaded', () => {
  fetch('/BookTalks-site/backend/reviews_list.php')
    .then(res => res.json())
    .then(data => {
      const container = document.querySelector('.resenz');
      container.innerHTML = '';
      data.reviews.forEach(r => {
        const block = document.createElement('div');
        block.className = 'resBook';
        let coverPath = r.coverimage_filename ? `/BookTalks-site/frontend/${r.coverimage_filename}` : '/BookTalks-site/uploads/no_cover.png';
        block.innerHTML = `
          <div class="book">
            <img src="${coverPath}" class="knizka" alt="Обложка книги" onerror="this.onerror=null;this.src='/BookTalks-site/uploads/no_cover.png'">
            <div class="name">
              <p class="nameBook">${r.book_name}</p>
              <p class="author">${r.author_first_name} ${r.author_last_name}</p>
            </div>
          </div>
          <div class="res">
            <div class="res-title">Рецензия</div>
            <hr class="res-line">
            <div class="res-disc">${r.text}</div>
            <div class="personName">${r.user_nickname}</div>
          </div>
        `;
        container.appendChild(block);
      });
    });
});
