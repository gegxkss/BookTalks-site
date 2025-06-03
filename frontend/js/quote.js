document.addEventListener('DOMContentLoaded', () => {
  fetch('/BookTalks-site/backend/quotes_list.php')
    .then(res => res.json())
    .then(data => {
      const container = document.querySelector('.resenz');
      container.innerHTML = '';
      data.quotes.forEach(q => {
        const block = document.createElement('div');
        block.className = 'resBook';
        let coverPath = q.coverimage_filename ? `/BookTalks-site/frontend/${q.coverimage_filename}` : '/BookTalks-site/uploads/no_cover.png';
        block.innerHTML = `
          <div class="book">
            <img src="${coverPath}" class="knizka" alt="Обложка книги" onerror="this.onerror=null;this.src='/BookTalks-site/uploads/no_cover.png'">
            <div class="name">
              <p class="nameBook">${q.book_name}</p>
              <p class="author">${q.author_first_name} ${q.author_last_name}</p>
            </div>
          </div>
          <div class="res">
            <div class="res-title">Цитата</div>
            <hr class="res-line">
            <div class="res-disc">${q.text}</div>
            <div class="personName">${q.user_nickname}</div>
          </div>
        `;
        container.appendChild(block);
      });
    });
});
