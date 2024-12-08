/* global use, db */
// MongoDB Playground
// To disable this template go to Settings | MongoDB | Use Default Template For Playground.
// Make sure you are connected to enable completions and to be able to run a playground.
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.
// The result of the last command run in a playground is shown on the results panel.
// By default the first 20 documents will be returned with a cursor.
// Use 'console.log()' to print to the debug output.
// For more documentation on playgrounds please refer to
// https://www.mongodb.com/docs/mongodb-vscode/playgrounds/

// Select the database to use.
use('mongodbVSCodePlaygroundDB');

// Insert a few documents into the sales collection.
db.getCollection('sales').insertMany([
    { 'item': 'abc', 'price': 10, 'quantity': 2, 'date': new Date('2014-03-01T08:00:00Z') },
    { 'item': 'jkl', 'price': 20, 'quantity': 1, 'date': new Date('2014-03-01T09:00:00Z') },
    { 'item': 'xyz', 'price': 5, 'quantity': 10, 'date': new Date('2014-03-15T09:00:00Z') },
    { 'item': 'xyz', 'price': 5, 'quantity': 20, 'date': new Date('2014-04-04T11:21:39.736Z') },
    { 'item': 'abc', 'price': 10, 'quantity': 10, 'date': new Date('2014-04-04T21:23:13.331Z') },
    { 'item': 'def', 'price': 7.5, 'quantity': 5, 'date': new Date('2015-06-04T05:08:13Z') },
    { 'item': 'def', 'price': 7.5, 'quantity': 10, 'date': new Date('2015-09-10T08:43:00Z') },
    { 'item': 'abc', 'price': 10, 'quantity': 5, 'date': new Date('2016-02-06T20:20:13Z') },
]);

// Run a find command to view items sold on April 4th, 2014.
const salesOnApril4th = db.getCollection('sales').find({
    date: { $gte: new Date('2014-04-04'), $lt: new Date('2014-04-05') }
}).count();

// Print a message to the output window.
console.log(`${salesOnApril4th} sales occurred in 2014.`);

// Here we run an aggregation and open a cursor to the results.
// Use '.toArray()' to exhaust the cursor to return the whole result set.
// You can use '.hasNext()/.next()' to iterate through the cursor page by page.
db.getCollection('sales').aggregate([
    // Find all of the sales that occurred in 2014.
    { $match: { date: { $gte: new Date('2014-01-01'), $lt: new Date('2015-01-01') } } },
    // Group the total sales for each product.
    { $group: { _id: '$item', totalSaleAmount: { $sum: { $multiply: ['$price', '$quantity'] } } } }
]);

// colection users
use('unairsatu_v2');
db.getCollection('users').updateMany({}, { $set: { role_aktif: 1 } });
function generateRandomString(length) {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
}

// rubah semua role_aktif pada users menjadi 1

db.getCollection('users').insertOne([
    {
        "id_role": ObjectId('674039236461fc1488d67fec'),
        "id_jenis_user": ObjectId('6740480562107a5c26a1f84c'),
        "username": "rakapurbayu",
        "nm_user": "rakapurbayu_123",
        "pass": "$2a$10$JVFps5uH7oMauLaKZJHdX.IuAbndwa5/CPwQJRRCC.DfgKEKMg4zC",
        "email": "raka@gmail.com",
        "role_aktif": ObjectId('674039236461fc1488d67fec'),
        "created_at": new Date(),
        "created_by": 1,
        "updated_at": new Date(),
        "updated_by": 1,
        "auth_key": generateRandomString(32),
        "jenis_kelamin": 1,
        "photo": "./storage/images/20241114160556.637.jpg",
        "phone": "085761875941",
        "token": generateRandomString(32),
        "pass_2": "$2a$10$JVFps5uH7oMauLaKZJHdX.IuAbndwa5/CPwQJRRCC.DfgKEKMg4zC",
        "moduls": [
            {
                "id_modul": "67403a6e19038f97f16e9984",
                "nm_modul": "Email Unair",
                "ket_modul": "Aplikasi Email Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
                "gbr_icon": "icon/email.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e9985",
                "nm_modul": "Cyber Campus",
                "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
                "gbr_icon": "icon/cybercampus.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e9989",
                "nm_modul": "E-Office",
                "ket_modul": "Aplikasi E-Office Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/eoffice.php",
                "gbr_icon": "icon/eoffice.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e998c",
                "nm_modul": "Dashboard",
                "ket_modul": "Aplikasi Dashboard Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dashboard.php",
                "gbr_icon": "icon/dashboard.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e998d",
                "nm_modul": "Simba",
                "ket_modul": "Aplikasi Simba Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/simba.php",
                "gbr_icon": "icon/simba.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e998a",
                "nm_modul": "VPN Unair",
                "ket_modul": "Aplikasi VPN Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
                "gbr_icon": "icon/vpn.png"
            },
            {
                "id_modul": "67403a6e19038f97f16e998b",
                "nm_modul": "Helpdesk Unair",
                "ket_modul": "Aplikasi Helpdesk Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/helpdesk.php",
                "gbr_icon": "icon/helpdesk.png"
            }
        ]
    }
]);

// update created at dan updated at pada users dengan object id 6754acadeb72108fb2ef5514
use('unairsatu_v2');
db.getCollection('users').updateOne({ _id: ObjectId('6754acadeb72108fb2ef5514') }, { $set: { created_at: new Date(), updated_at: new Date() } });

db.getCollection('users').insertMany([
    {
        "id_role": ObjectId('674039236461fc1488d67fed'),
        "id_jenis_user": ObjectId('6740480562107a5c26a1f84c'),
        "username": "03ghandi.mahasiswa",
        "nm_user": "Ghandi Nur Mahasiswa",
        "pass": "02ghandi@1003",
        "email": "muhamadghandinursetiawan@gmail.com",
        "role_aktif": ObjectId('674039236461fc1488d67fed'),
        "created_at": new Date(),
        "created_by": 1,
        "updated_at": new Date(),
        "updated_by": 1,
        "auth_key": generateRandomString(32),
        "jenis_kelamin": 1,
        "photo": "mahasiswa/ghandi.jpg",
        "phone": "085856490085",
        "token": generateRandomString(32),
        "pass_2": "02ghandi@1003",
        "moduls": [
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
                "nm_modul": "Email Unair",
                "ket_modul": "Aplikasi Email Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
                "gbr_icon": "icon/email.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
                "nm_modul": "Cyber Campus",
                "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
                "gbr_icon": "icon/cybercampus.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9986'),  // aplikasi pusba elpt
                "nm_modul": "Pusba ELPT",
                "ket_modul": "Aplikasi Pusba ELPT Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/pusba.php",
                "gbr_icon": "icon/pusba.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
                "nm_modul": "VPN Unair",
                "ket_modul": "Aplikasi VPN Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
                "gbr_icon": "icon/vpn.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
                "nm_modul": "Helpdesk Unair",
                "ket_modul": "Aplikasi Helpdesk Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/helpdesk.php",
                "gbr_icon": "icon/helpdesk.png"
            }
        ]
    },
    {
        "id_role": ObjectId('674039236461fc1488d67fed'),
        "id_jenis_user": ObjectId('6740480562107a5c26a1f84d'),
        "username": "alifian_sukma",
        "nm_user": "alifian",
        "pass": "10alifian.45",
        "email": "alifiansukma01@gmail.com",
        "role_aktif": ObjectId('674039236461fc1488d67fed'),
        "created_at": new Date(),
        "created_by": 1,
        "updated_at": new Date(),
        "updated_by": 1,
        "auth_key": generateRandomString(32),
        "jenis_kelamin": 1,
        "photo": "dosen/alifian.jpg",
        "phone": "087475896142",
        "token": generateRandomString(32),
        "pass_2": "10alifian.45",
        "moduls": [
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
                "nm_modul": "Email Unair",
                "ket_modul": "Aplikasi Email Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
                "gbr_icon": "icon/email.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
                "nm_modul": "Cyber Campus",
                "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
                "gbr_icon": "icon/cybercampus.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9987'),  // aplikasi dosen
                "nm_modul": "Dosen",
                "ket_modul": "Aplikasi Dosen Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dosen.php",
                "gbr_icon": "icon/dosen.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9988'),  // aplikasi dosen v3
                "nm_modul": "Dosen V3",
                "ket_modul": "Aplikasi Dosen V3 Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dosen_v3.php",
                "gbr_icon": "icon/dosen_v3.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
                "nm_modul": "Dashboard",
                "ket_modul": "Aplikasi Dashboard Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dashboard.php",
                "gbr_icon": "icon/dashboard.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
                "nm_modul": "VPN Unair",
                "ket_modul": "Aplikasi VPN Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
                "gbr_icon": "icon/vpn.png"
            }
        ]
    },
    {
        "id_role": ObjectId('674039236461fc1488d67fed'),
        "id_jenis_user": ObjectId('6740480562107a5c26a1f84e'),
        "username": "01Fiki.tendik",
        "nm_user": "fiki tendik",
        "pass": "fiki@.1987",
        "email": "fikitendikunair@gmail.com",
        "role_aktif": ObjectId('674039236461fc1488d67fed'),
        "created_at": new Date(),
        "created_by": 1,
        "updated_at": new Date(),
        "updated_by": 1,
        "auth_key": generateRandomString(32),
        "jenis_kelamin": 1,
        "photo": "tendik/fiki.jpg",
        "phone": "082412632478",
        "token": generateRandomString(32),
        "pass_2": "fiki@.1987",
        "moduls": [
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
                "nm_modul": "Email Unair",
                "ket_modul": "Aplikasi Email Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
                "gbr_icon": "icon/email.png"
            },
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
                "nm_modul": "Cyber Campus",
                "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
                "gbr_icon": "icon/cybercampus.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
                "nm_modul": "E-Office",
                "ket_modul": "Aplikasi E-Office Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/eoffice.php",
                "gbr_icon": "icon/eoffice.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
                "nm_modul": "Dashboard",
                "ket_modul": "Aplikasi Dashboard Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dashboard.php",
                "gbr_icon": "icon/dashboard.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
                "nm_modul": "VPN Unair",
                "ket_modul": "Aplikasi VPN Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
                "gbr_icon": "icon/vpn.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
                "nm_modul": "Helpdesk Unair",
                "ket_modul": "Aplikasi Helpdesk Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/helpdesk.php",
                "gbr_icon": "icon/helpdesk.png"
            }
        ]
    },
    {
        "id_role": ObjectId('674039236461fc1488d67fed'),
        "id_jenis_user": ObjectId('6740480562107a5c26a1f84f'),
        "username": "03fitri.kps",
        "nm_user": "fitri kps",
        "pass": "kps.fitri02@kps",
        "email": "fikitendikunair@gmail.com",
        "role_aktif": ObjectId('674039236461fc1488d67fed'),
        "created_at": new Date(),
        "created_by": 1,
        "updated_at": new Date(),
        "updated_by": 1,
        "auth_key": generateRandomString(32),
        "jenis_kelamin": 1,
        "photo": "kps/fitri.jpg",
        "phone": "082412632478",
        "token": generateRandomString(32),
        "pass_2": "kps.fitri02@kps",
        "moduls": [
            {
                "modul_id": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
                "nm_modul": "Email Unair",
                "ket_modul": "Aplikasi Email Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
                "gbr_icon": "icon/email.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
                "nm_modul": "Cyber Campus",
                "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
                "gbr_icon": "icon/cybercampus.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
                "nm_modul": "E-Office",
                "ket_modul": "Aplikasi E-Office Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/eoffice.php",
                "gbr_icon": "icon/eoffice.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
                "nm_modul": "Dashboard",
                "ket_modul": "Aplikasi Dashboard Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/dashboard.php",
                "gbr_icon": "icon/dashboard.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
                "nm_modul": "VPN Unair",
                "ket_modul": "Aplikasi VPN Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
                "gbr_icon": "icon/vpn.png"
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
                "nm_modul": "Helpdesk Unair",
                "ket_modul": "Aplikasi Helpdesk Universitas Airlangga",
                "alamat": "http://cybercampus.unair.ac.id//sso/helpdesk.php",
                "gbr_icon": "icon/helpdesk.png"
            }
        ]
    }
]);

// collection jenis users
use('unairsatu_v2');
db.getCollection('jenis_user').insertMany([
    {
        "nm_jenis_user": "Mahasiswa",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9986'),  // aplikasi pusba elpt
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    },
    {
        "nm_jenis_user": "Dosen",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9987'),  // aplikasi dosen
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9988'),  // aplikasi dosen v3
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
        ]
    },
    {
        "nm_jenis_user": "Tendik",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    },
    {
        "nm_jenis_user": "KPS",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    },
    {
        "nm_jenis_user": "Dekanat",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998d'),  // aplikasi simba
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    },
    {
        "nm_jenis_user": "Ketua_Unit",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998d'),  // aplikasi simba
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    },
    {
        "nm_jenis_user": "Pimpinan_univ",
        "templates": [
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9984'),  // aplikasi email unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9985'),  // aplikasi cyber campus
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e9989'),  // aplikasi e-office
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998c'),  // aplikasi dashboard
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998d'),  // aplikasi simba
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998a'),  // aplikasi vpn unair
            },
            {
                "id_modul": ObjectId('67403a6e19038f97f16e998b'),  // aplikasi helpdesk unair
            }
        ]
    }
]);

// colection modul
use('unairsatu_v2');
db.getCollection('modul').insertMany([
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "Email Unair",
        "ket_modul": "Aplikasi Email Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/email.php",
        "urutan": 1,
        "gbr_icon": "icon/email.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/email.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "Cyber Campus",
        "ket_modul": "Aplikasi Cyber Campus Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/cybercampus.php",
        "urutan": 2,
        "gbr_icon": "icon/cybercampus.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/cybercampus.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "Pusba ELPT",
        "ket_modul": "Aplikasi Pusba ELPT Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/pusba.php",
        "urutan": 3,
        "gbr_icon": "icon/pusba.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/pusba.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "Dosen",
        "ket_modul": "Aplikasi Dosen Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/dosen.php",
        "urutan": 4,
        "gbr_icon": "icon/dosen.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/dosen.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "Dosen V3",
        "ket_modul": "Aplikasi Dosen V3 Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/dosen_v3.php",
        "urutan": 5,
        "gbr_icon": "icon/dosen_v3.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/dosen_v3.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53af'),  // Kategori Aplikasi Kampus
        "nm_modul": "E-Office",
        "ket_modul": "Aplikasi E-Office Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/eoffice.php",
        "urutan": 6,
        "gbr_icon": "icon/eoffice.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/eoffice.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b0'),  // kategori Aplikasi Layanan Kampus
        "nm_modul": "VPN Unair",
        "ket_modul": "Aplikasi VPN Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/vpn.php",
        "urutan": 1,
        "gbr_icon": "icon/vpn.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/vpn.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b0'),  // kategori Aplikasi Layanan Kampus
        "nm_modul": "Helpdesk Unair",
        "ket_modul": "Aplikasi Helpdesk Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/helpdesk.php",
        "urutan": 2,
        "gbr_icon": "icon/helpdesk.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/helpdesk.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b1'),  // kategori Aplikasi Dashboard Kampus
        "nm_modul": "Dashboard",
        "ket_modul": "Aplikasi Dashboard Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/dashboard.php",
        "urutan": 1,
        "gbr_icon": "icon/dashboard.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/dashboard.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b2'),  // kategori Airlangga Resource Planning
        "nm_modul": "Simba",
        "ket_modul": "Aplikasi Simba Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/simba.php",
        "urutan": 1,
        "gbr_icon": "icon/simba.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/simba.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b3'),  // kategori Aplikasi khusus
        "nm_modul": "E-Lab",
        "ket_modul": "Aplikasi untuk mengakses dan memesan layanan laboratorium di Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/elab.php",
        "urutan": 1,
        "gbr_icon": "icon/elab.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/elab.png"
    },
    {
        "id_kategori": ObjectId('674036a5712f74225a5f53b3'),  // kategori Aplikasi khusus
        "nm_modul": "Dashboard Admin",
        "ket_modul": "Aplikasi Dashboard Admin Universitas Airlangga",
        "is_aktif": "1",
        "alamat": "http://cybercampus.unair.ac.id//sso/dashboard_admin.php",
        "urutan": 1,
        "gbr_icon": "icon/dashboard_admin.png",
        "created_at": new Date(),
        "created_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "updated_at": new Date(),
        "updated_by": ObjectId('66fef5722cc0de4b6829e8db'),
        "icon": "icon_kc/dashboard_admin.png"
    }
]);

// colection role
use('unairsatu_v2');
db.getCollection('role').updateMany(
    {},
    {
        $set: {
            created_by: ObjectId('674039236461fc1488d67fec'),
            updated_by: ObjectId('674039236461fc1488d67fec')
        }
    }
);
use('unairsatu_v2');
db.getCollection('role').insertMany([
    {
        "nm_role": "Admin",
        created_at: new Date(),
        created_by: 1,
        updated_at: new Date(),
        updated_by: 1
    },
    {
        "nm_role": "civitas",
        created_at: new Date(),
        created_by: 1,
        updated_at: new Date(),
        updated_by: 1
    }
]);

// colection kategori
use('unairsatu_v2');
db.getCollection('kategori').insertMany([
    {
        "nm_kategori": "Aplikasi Kampus"
    },
    {
        "nm_kategori": "Aplikasi Layanan Kampus"
    },
    {
        "nm_kategori": "Aplikasi Dashboard Kampus"
    },
    {
        "nm_kategori": "Airlangga Resource Planning"
    },
    {
        "nm_kategori": "Aplikasi Khusus"
    }
]);