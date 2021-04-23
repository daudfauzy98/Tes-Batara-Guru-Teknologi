# REST API Pokemon

REST API ini mengambil data daftar pokemon dari API http://pokeapi.co/ yang sudah berbentuk JSON, kemudian disimpan ke dalam database MongoDB. API ini memiliki batas request sebanyak 5 request permenit. Jika lebih dari 5 request, client akan menunggu 1 menit untuk melihat hasil request-nya. Selain itu, untuk mengakses API ini diperlukan autentikasi berupa Header Authorization dengan value "izin-masuk-gan". Screenshot hasil bisa dilihat di menu Issues.
