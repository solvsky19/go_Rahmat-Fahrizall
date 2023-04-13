- 
Hari ini saya belajar tentang Docker, Docker adalah platform perangkat lunak yang memungkinkan pengembang untuk mengemas aplikasi mereka dan dependensi mereka ke dalam kontainer yang dapat dijalankan di mana saja. Kontainer adalah lingkungan runtime yang independen yang terisolasi dari lingkungan host dan kontainer lainnya, tetapi masih dapat berinteraksi dengan host dan kontainer lainnya melalui jaringan yang terpisah.

    - Virtual Machine adalah lingkungan virtualisasi lengkap yang memungkinkan sistem operasi dan aplikasi dijalankan pada level yang sangat terisolasi, terpisah dari lingkungan host. Sedangkan Container hanya mengisolasi aplikasi dan dependensinya, tetapi berbagi kernel dan beberapa sumber daya lainnya dengan lingkungan host.

    - Virtual Machine membutuhkan hypervisor untuk menjalankan sistem operasi dan aplikasi pada lingkungan virtualnya. Hypervisor ini memerlukan sumber daya yang cukup besar, sehingga membutuhkan lebih banyak sumber daya dari host. Sedangkan Container tidak memerlukan hypervisor dan memerlukan sumber daya yang lebih sedikit dibandingkan dengan VM.

    - Virtual Machine memerlukan sistem operasi yang terpisah untuk setiap VM, sehingga membutuhkan penyimpanan yang lebih besar dan waktu instalasi yang lebih lama. Sedangkan Container berbagi kernel dengan lingkungan host dan hanya memerlukan sedikit overhead untuk membuat dan menjalankan kontainer.

    - Virtual Machine lebih cocok untuk menjalankan aplikasi dan sistem yang membutuhkan isolasi dan keamanan yang tinggi, serta ketahanan terhadap gangguan atau serangan dari lingkungan host. Sedangkan Container lebih cocok untuk menjalankan aplikasi yang ringan dan terisolasi, yang dapat dengan mudah diatur, diperbarui, dan diintegrasikan dengan lingkungan yang lebih besar.

    - `RUN`     : Perintah untuk membuat dan menjalankan kontainer Docker dari sebuah image.
    - `PULL`    : Perintah untuk mengunduh sebuah image Docker dari registry.
    - `PUSH`    : Perintah untuk mengirimkan sebuah image Docker ke registry.
    - `BUILD`   : Perintah untuk membuat sebuah image Docker dari Dockerfile.
    - `IMAGES`  : Perintah untuk menampilkan daftar image Docker yang tersimpan di host.
    - `PS`      : Perintah untuk menampilkan daftar kontainer Docker yang sedang berjalan.
    - `START`   : Perintah untuk memulai sebuah kontainer Docker yang sudah dibuat.
    - `STOP`    : Perintah untuk menghentikan sebuah kontainer Docker yang sedang berjalan.
    - `RM`      : Perintah untuk menghapus sebuah kontainer Docker.
    - `RMI`     : Perintah untuk menghapus sebuah image Docker.
    - `EXEC`    : Perintah untuk mengeksekusi perintah di dalam sebuah kontainer Docker yang sedang berjalan.
    - `NETWORK` : Perintah untuk mengelola jaringan Docker.
    - `COMPOSE` : Perintah untuk mengelola aplikasi Docker yang kompleks dengan menggunakan file konfigurasi YAML.

