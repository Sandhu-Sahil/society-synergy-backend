package services

import "fmt"

func generateHtmlContent(otp string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Congratulations!</title>
		<style>
			/* Inline CSS for styling */
			.out {
				font-family: Arial, sans-serif;
				background-color: #000000;
				padding: 10px;
			}
            .main{
                text-align: center;
            }
			.logo {
				width: 100px;
				height: 100px;
                text-align: center;
			}
			.congratulations {
				font-size: 18px;
				margin-top: 20px;
                color: #ffffff;
                text-align: center;
			}
			.otp {
				font-size: 24px;
				font-weight: bold;
				margin-top: 10px;
                color: #ffffff;
                text-align: center;
			}
			.copy-button {
				background-color: #ff6a00;
				color: #000000;
				border: none;
				padding: 5px 10px;
				border-radius: 5px;
				cursor: pointer;
				margin-top: 10px;
                text-align: center;
			}
			.copy-button:hover {
				background-color: #0056b3;
                text-align: center;
			}
            
			.website-link {
				margin-top: 20px;
				text-decoration: none;
				color: #ff6a00;
                text-align: center;
			}
			.thank-you {
				font-size: 14px;
				margin-top: 20px;
                color: #ffffff;
                background-color: #1b1b1b;
                padding: 15px;
                margin: 30px;
			}

            .otp-image {
                width: 250px;
            }
		</style>
	</head>
	<body>
    <div class="out">
		<!-- Logo in SVG format -->
        <div class="main">
		<img src="https://cdn.discordapp.com/attachments/888075387228278817/1157367313675333692/logo.png?ex=651859ce&is=6517084e&hm=66401bbae6be874a7711b20b7602e0bbfbd12a95c39a01519ee2f0bfb1bb6877&" alt="text" width="80" border="0">
        <br>
        <img src="https://cdn.discordapp.com/attachments/888075387228278817/1157359552401002646/text.png?ex=65185293&is=65170113&hm=ffb3837e1f53676fe9e9f461d3bbba4bcde077cd6d294ecb4fe3c30346b27bd1&" alt="text" border="0" class="otp-image">
	
		<!-- Congratulations message -->
		<div class="congratulations">Congratulations! You've successfully registered on Society Synergy.</div>
	
		<!-- OTP -->
		<div class="otp">Your OTP is: <strong><span id="otp">%s</span></strong></div>
		<button class="copy-button" onclick="
            navigator.clipboard.writeText('%s');
        ">Copy OTP</button>
	
		<!-- Link to your website -->
        <br><br>
		<a class="website-link" href="#">Visit Our Website</a>
        </div>
	
		<!-- Thank you message -->
		<div class="thank-you">Thank you for registering, Team Society Synergy<br><hr><br>Sahilsher Singh Sandhu<br>JIIT sector-128 <br>@2025<br>9921103131@mail.jiit.ac.in <br>sandhu.sahil2002@gmail.com</div>
    </div>
	</body>
	</html>
	`, otp, otp)
}
