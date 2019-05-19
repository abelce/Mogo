{{define "footer"}}
<div id="footer">
 	<div id="footer-inner">
 		<p id="copyright">Copyright (c) {{"copyright.beginYear" | get}} - {{"copyright.endYear" | get}} {{"copyright.owner"|get}} &nbsp;
 		Powered by <a href="https://github.com/scottkiss/gosk">gosk</a>
 		</p>
 		</div>  
 	</div>
	<script type="text/javascript" src="/assets/themes/{{"theme"|get}}/plugin/prettify/prettify.js"></script>
 </body>
 </html>
{{end}}