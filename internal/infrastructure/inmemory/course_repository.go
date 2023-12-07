package inmemory

import (
	"learn-to-code/internal/domain/quiz/course"
)

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

// CourseRepository contains hardcoded data for now to validate the requirements and access patterns
type CourseRepository struct {
}

const CourseID = "ed86d338-84a0-4486-a314-b99b17175875"
const CourseStepID = "c7486278-a50c-4629-89b9-cc1c74d7a538"
const QuizID = "fcf7890f-9c72-46d3-931e-34494307be37"
const FirstQuestionID = "14c20d31-c7e1-416d-9c8e-1f2040141f0b"
const FirstAnswerID = "06a1956e-b659-493f-9533-b27733ddd7fe"
const FirstCorrectAnswerID = "48a293ee-7f43-4e3d-85d1-4737e6385c7c"

func (q *CourseRepository) FindByID(id string) (course.Course, error) {

	if id == CourseID {
		return course.Course{
			ID:   CourseID,
			Name: "Frontend Development",
			Steps: []course.Step{
				{
					ID:   CourseStepID,
					Name: "The essentials of the Web",
					Quizzes: []course.StepQuiz{
						{
							ID: QuizID,
							Questions: []course.QuizQuestion{
								{
									ID:          FirstQuestionID,
									Text:        "What is HTML used for?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          FirstAnswerID,
											Text:        "Styling the website",
											IsCorrect:   false,
											Description: "Styling a website is primarily done using CSS, which helps in determining colors, font sizes, positioning, and other visual elements.",
										},
										{
											ID:          FirstCorrectAnswerID,
											Text:        "Describing the structure of a webpage",
											IsCorrect:   true,
											Description: "HTML stands for HyperText Markup Language and it provides the basic structure to a webpage. Through HTML, the browser understands headings, paragraphs, links, and other content types.",
										},
										{
											ID:          "0aaf5a53-ec6b-4487-962c-e22f8f2ce45e",
											Text:        "Creating animations for a webpage",
											IsCorrect:   false,
											Description: "While animations can be added to a webpage, they are typically achieved using both CSS and JavaScript, depending on the complexity. HTML itself isn't used for animations.",
										},
										{
											ID:          "870cc720-b6e6-447b-afee-09b1b44b7730",
											Text:        "Accessing web databases",
											IsCorrect:   false,
											Description: "Accessing databases is generally the realm of backend languages and APIs, not directly through HTML.",
										},
									},
								},
								{
									ID:          "343f51b8-68b3-4248-be34-981e4040d892",
									Text:        "Which of the following is NOT an HTML element?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "289016f8-d546-489b-b50d-5b9193cda0f4",
											Text:        "<heading>",
											IsCorrect:   false,
											Description: "The term <heading> is not used as an HTML element. Instead, we have heading tags that range from <h1> to <h6> to denote headings of different levels.",
										},
										{
											ID:          "1c2f637a-be46-48ef-84b0-74b8dd97f550",
											Text:        "<paragraph>",
											IsCorrect:   false,
											Description: "<paragraph> is not the correct HTML element for creating paragraphs. The right tag for a paragraph in HTML is <p>.",
										},
										{
											ID:          "2473afe5-bfc7-4903-8f67-c8e109860420",
											Text:        "<list>",
											IsCorrect:   false,
											Description: "While you can create lists in HTML, the tag <list> doesn't exist. For lists, we use <ul> for unordered lists and <ol> for ordered lists.",
										},
										{
											ID:          "846326cd-d809-496e-9789-7588401101ce",
											Text:        "None of the above",
											IsCorrect:   true,
											Description: "All of the provided options are not standard HTML elements.",
										},
									},
								},
								{
									ID:          "12b39c35-aebe-48ad-8125-6f93f73d6bbd",
									Text:        "Which language is primarily used to determine the visual style of a website?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "925a33b4-7a3e-4f1d-a4ab-259c0e4807ff",
											Text:        "JavaScript",
											IsCorrect:   false,
											Description: "JavaScript is mainly used for adding functionality and interactivity to web pages.",
										},
										{
											ID:          "92b56418-e0c9-4b3b-864a-f39e639fcf55",
											Text:        "Python",
											IsCorrect:   false,
											Description: "Python is a programming language that's often used for backend web development, data analysis, artificial intelligence, and more.",
										},
										{
											ID:          "38cabc9b-8a88-405e-b21a-dff76bed9799",
											Text:        "HTML",
											IsCorrect:   false,
											Description: "HTML is the markup language used to structure content on the web. While it can carry some style attributes, CSS is the primary tool for styling.",
										},
										{
											ID:          "6cc58e72-8d1f-4cc2-9252-b5dd60e4463f",
											Text:        "CSS",
											IsCorrect:   true,
											Description: "CSS (Cascading Style Sheets) is the language specifically designed to style web content. It defines how elements should appear in terms of layout, colors, fonts, and more.",
										},
									},
								},
								{
									ID:          "094a9c58-b252-4019-8d1a-69766352a773",
									Text:        "What is the primary purpose of visual design in web development?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "b83d7bbb-20b2-4998-b50f-9260b615a67d",
											Text:        "Accelerating website speed",
											IsCorrect:   false,
											Description: "While a well-designed website can improve user experience, the speed of a website is more about optimization, efficient coding, and hosting solutions.",
										},
										{
											ID:          "db866e26-9027-4fa4-891b-955e12524612",
											Text:        "Ensuring website security",
											IsCorrect:   false,
											Description: "Website security is crucial, but visual design isn't directly related to it. Security measures involve both frontend and backend strategies, encryptions, and more.",
										},
										{
											ID:          "573d00b0-aa72-4bf4-98a7-0098722ecc81",
											Text:        "Delivering a unique message and enhancing user experience",
											IsCorrect:   true,
											Description: "Visual design uses typography, color, graphics, and more to help convey a unique message and provide a memorable user experience.",
										},
										{
											ID:          "46eda80d-7bec-4b10-9b96-bd506f2acffd",
											Text:        "Backing up website data",
											IsCorrect:   false,
											Description: "Backing up website data is an administrative task and not a direct responsibility of visual design.",
										},
									},
								},
								{
									ID:          "a5d7ccec-a8bc-478b-a9b3-0d6493802408",
									Text:        "Why is web accessibility important?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "d3c8b411-626e-48e8-b414-f123093cab5c",
											Text:        "It ensures that websites look the same on all devices.",
											IsCorrect:   false,
											Description: "While a consistent appearance across devices is essential, that's more about responsive design than accessibility.",
										},
										{
											ID:          "834e9bea-947e-4ae4-9d1a-fd5fd08bae36",
											Text:        "It makes websites load faster.",
											IsCorrect:   false,
											Description: "Web accessibility doesn't necessarily mean faster load times. Load times depend on optimization techniques, server speeds, and more.",
										},
										{
											ID:          "85b57b14-7349-4c0f-b21d-4b92bcdc1507",
											Text:        "It guarantees that websites can be understood and navigated by everyone, including people with disabilities.",
											IsCorrect:   true,
											Description: "Accessibility ensures that web content and user interfaces are accessible to all users, regardless of physical or cognitive disabilities, ensuring inclusivity on the web.",
										},
										{
											ID:          "506033f3-6c24-4837-afaf-2c75e55f3283",
											Text:        "It ensures higher search engine rankings.",
											IsCorrect:   false,
											Description: "While search engines do value accessible websites, the primary goal of web accessibility is inclusiveness, not SEO.",
										},
									},
								},
								{
									ID:          "4fa7cd1d-616e-47db-be8e-3f9fc3720dff",
									Text:        "What does responsive web design ensure?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "48ce094c-0b9a-455f-9b4c-2931252eeb05",
											Text:        "Faster loading speed on mobile devices",
											IsCorrect:   false,
											Description: "Responsive design doesn't inherently speed up website loading times. Its main goal is adaptability.",
										},
										{
											ID:          "1631d3b5-0e30-4bb2-a23f-41b7206b1731",
											Text:        "Higher security against web threats",
											IsCorrect:   false,
											Description: "Responsive design focuses on layout adaptability, not security.",
										},
										{
											ID:          "e9f67fda-0eb2-4647-ae9e-eae4aa4c654e",
											Text:        "Flexibility in website appearance according to device size and orientation",
											IsCorrect:   true,
											Description: "Responsive web design allows websites to adjust and look good on devices of all sizes, whether it's a large monitor, tablet, or smartphone.",
										},
										{
											ID:          "d7b900e1-d4d0-46a5-b1a7-8a35393eb666",
											Text:        "Vibrant color schemes for websites",
											IsCorrect:   false,
											Description: "Color schemes are part of visual design. Responsiveness is about layout adjustments based on screen sizes.",
										},
									},
								},
								{
									ID:          "39dc2763-2abb-4618-98f4-9634b33a497a",
									Text:        "Which of the following is a powerful layout method introduced in CSS3?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "b288f157-9216-4103-8eaa-6c88f6d3dbcb",
											Text:        "CSS Shapes",
											IsCorrect:   false,
											Description: "CSS Shapes is about wrapping content around custom paths, not about layout per se.",
										},
										{
											ID:          "aedab56a-25ec-45b8-90af-b3a1b4653ad2",
											Text:        "CSS Flexbox",
											IsCorrect:   true,
											Description: "CSS Flexbox provides an efficient way to lay out, align, and distribute space among items in a container, even when their sizes are unknown or dynamic.",
										},
										{
											ID:          "60a18ebc-56a6-4ffa-98b9-5cb663a022df",
											Text:        "CSS Templates",
											IsCorrect:   false,
											Description: "There's no standard feature named 'CSS Templates'.",
										},
										{
											ID:          "81d98e55-0e79-4db1-bd78-94ec5fc102ef",
											Text:        "CSS Fonts",
											IsCorrect:   false,
											Description: "CSS Fonts relates to font-face and loading custom fonts, not layouts.",
										},
									},
								},
								{
									ID:          "305ac6e1-7c85-42b8-ae4c-ef24a9f281d0",
									Text:        "What does the CSS Grid enable in terms of layout design?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "01a107ce-e2fd-448d-81e0-dceb2be37aac",
											Text:        "It allows embedding multimedia.",
											IsCorrect:   false,
											Description: "While multimedia can be placed within a CSS Grid, the grid itself is not specifically for embedding multimedia.",
										},
										{
											ID:          "78beafed-7bf6-46e7-b70e-9b93b5eb9eb3",
											Text:        "It ensures website security.",
											IsCorrect:   false,
											Description: "CSS Grid focuses on layout and has nothing to do with security.",
										},
										{
											ID:          "55a45372-4f55-4aa8-9db4-c49387f7ea1a",
											Text:        "It enables easy creation of flexible, two-dimensional layouts.",
											IsCorrect:   true,
											Description: "CSS Grid provides a method for placing elements within columns and rows, making complex responsive layouts more manageable.",
										},
										{
											ID:          "e7fa6ca5-b72f-4d06-af7c-26d21cb4b64b",
											Text:        "It automates browser compatibility.",
											IsCorrect:   false,
											Description: "Browser compatibility depends on various factors. While Grid is supported in many modern browsers, its use doesn't 'automate' compatibility.",
										},
									},
								},
								{
									ID:          "8ffe3cda-9b9f-4aac-8813-3f0c73668166",
									Text:        "Which of the following elements can describe text as a heading in HTML?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "8c2f89eb-fc21-49cc-adc4-8ed330d4c5a3",
											Text:        "<header>",
											IsCorrect:   false,
											Description: "The <header> element typically contains introductory content or navigation links and not used as a standard heading.",
										},
										{
											ID:          "58a7b8e4-638a-49a3-8809-0f2513fa3093",
											Text:        "<heading>",
											IsCorrect:   false,
											Description: "There is no standard <heading> tag in HTML.",
										},
										{
											ID:          "bc81de7d-8fd4-4f8f-9b3d-6abae72fb020",
											Text:        "<h1>",
											IsCorrect:   true,
											Description: "The <h1> tag in HTML is used to represent the top-level heading, making it the most important among the heading tags ranging from <h1> to <h6>.",
										},
										{
											ID:          "48fd0307-1982-431b-b188-60d0d19b1b73",
											Text:        "<top>",
											IsCorrect:   false,
											Description: "The <top> tag doesn't exist in standard HTML.",
										},
									},
								},
								{
									ID:          "4c1bc3d4-0c62-4299-895b-6ff826675d74",
									Text:        "When you want to make a list of items without any specific order, which HTML tag would you use?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "e0e13029-1ce7-40da-8e1c-c25b702ae20a",
											Text:        "<ol>",
											IsCorrect:   false,
											Description: "The <ol> tag is used for ordered lists, where the order of items is important.",
										},
										{
											ID:          "13d5072c-6b28-4008-afe6-5adf1d67fe3e",
											Text:        "<dl>",
											IsCorrect:   false,
											Description: "The <dl> tag is for description lists and isn't typically used for general lists.",
										},
										{
											ID:          "f055cf1e-78d2-4031-89ce-ef3fd28cc8ab",
											Text:        "<ul>",
											IsCorrect:   true,
											Description: "The <ul> tag stands for unordered list, which is used for lists where order doesn't matter.",
										},
										{
											ID:          "04fb764e-bcbb-42ea-a1e7-a49ccddd9918",
											Text:        "<list>",
											IsCorrect:   false,
											Description: "There is no standard <list> tag in HTML.",
										},
									},
								},
								{
									ID:          "62141a7b-e0ca-4862-bafa-beea085b13a6",
									Text:        "Considering the 'Cascading' in Cascading Style Sheets (CSS), which of the following statements is true?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "416c5722-a35a-4d9a-99d2-dec370711046",
											Text:        "External stylesheets have a higher priority than inline styles.",
											IsCorrect:   false,
											Description: "Inline styles have a higher priority than external stylesheets.",
										},
										{
											ID:          "6058c295-eeae-4ba2-960c-d132b3526dc9",
											Text:        "The last rule defined overrides any previous, conflicting rules.",
											IsCorrect:   true,
											Description: "Cascading means that the order of CSS rules matters. If two rules conflict, the latter rule will take precedence, given they have the same specificity.",
										},
										{
											ID:          "545feb47-ddf9-4e56-9345-71f32edf9ad8",
											Text:        "All styles have an equal weight, regardless of where they are defined.",
											IsCorrect:   false,
											Description: "Styles do not have equal weight; their application is determined by specificity and order.",
										},
										{
											ID:          "847fa93b-caad-4ab4-818a-13ad92297740",
											Text:        "Inline styles only apply to JavaScript-generated content.",
											IsCorrect:   false,
											Description: "Inline styles can be applied directly to any HTML element, not just JavaScript-generated content.",
										},
									},
								},
								{
									ID:          "63894dd1-1218-4db5-9da5-249995a1a4ab",
									Text:        "In responsive design, what does the 'viewport' refer to?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "ff9f0f84-5b73-4dde-abe6-87ce8ee300e1",
											Text:        "The visible area of a webpage that the user can interact with.",
											IsCorrect:   true,
											Description: "The viewport is the user's visible area of a web page. It varies with the device, and it can be controlled using the <meta> viewport element.",
										},
										{
											ID:          "b9902d50-b70f-466a-b804-29a93d3acabc",
											Text:        "The server's perspective on how to serve content.",
											IsCorrect:   false,
											Description: "The server doesn't have a 'perspective'; it serves content based on requests.",
										},
										{
											ID:          "d9ce9405-60bf-45b5-8055-e250207d4712",
											Text:        "The total area of a website, including off-screen content.",
											IsCorrect:   false,
											Description: "The viewport only concerns what's currently visible and does not include off-screen content.",
										},
										{
											ID:          "17ffd2f4-8494-4a9d-a066-f53e7150115a",
											Text:        "The framework used to build mobile-responsive designs.",
											IsCorrect:   false,
											Description: "While certain frameworks can help with responsive design, the term 'viewport' does not pertain to any specific framework.",
										},
									},
								},
								{
									ID:          "187c0ae5-1758-4388-9080-349fec8075ec",
									Text:        "Given the growth of mobile web browsing, why is Flexbox considered an important tool in responsive design?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "4fe7134c-ab89-4042-84a7-313b33936d7e",
											Text:        "It enables developers to use less JavaScript.",
											IsCorrect:   false,
											Description: "While Flexbox can reduce the need for some layout-related JavaScript, its main purpose isn't about reducing JavaScript usage.",
										},
										{
											ID:          "e5fb7cd0-7fe2-48ca-988f-cf4a82f48963",
											Text:        "It allows more vibrant color schemes.",
											IsCorrect:   false,
											Description: "Flexbox is about layout, not color schemes.",
										},
										{
											ID:          "ce1e82bb-54bd-4386-af1f-9a785c225d8c",
											Text:        "Flexbox layouts adapt to different screen sizes without requiring pixel-based dimensions.",
											IsCorrect:   true,
											Description: "Flexbox allows for flexible layouts that can adapt to various screen sizes and orientations, making it easier to create responsive designs without relying heavily on fixed dimensions.",
										},
										{
											ID:          "93841442-0a10-4334-8a03-7cc9ff1671df",
											Text:        "It's a plugin that makes websites mobile-friendly automatically.",
											IsCorrect:   false,
											Description: "Flexbox is not a plugin; it's a CSS layout model.",
										},
									},
								},
								{
									ID:          "45a339e2-6b49-4be2-a5d3-d85f9ffbf0af",
									Text:        "In the context of CSS Grid, what does the 'fr' unit stand for?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "29b649e7-e623-4118-8acc-d13e5f31e916",
											Text:        "Frame",
											IsCorrect:   false,
											Description: "While 'frame' sounds related to grids or layouts, it's not what 'fr' represents.",
										},
										{
											ID:          "ff424c40-1cc2-48a6-ba60-c4244f1ede46",
											Text:        "Fraction",
											IsCorrect:   true,
											Description: "The 'fr' unit in CSS Grid stands for 'fraction'. It represents a fraction of the available space in the grid container.",
										},
										{
											ID:          "8994fea8-c493-4ed9-a451-8b9a701f8f6f",
											Text:        "Frequency",
											IsCorrect:   false,
											Description: "Frequency is not a unit of measurement in CSS Grid.",
										},
										{
											ID:          "00e23b58-d846-40a3-a843-87eecebd6abf",
											Text:        "Format",
											IsCorrect:   false,
											Description: "Format doesn't correlate with the purpose of 'fr' in grid layouts.",
										},
									},
								},
								{
									ID:          "621ebc95-6598-49e2-98e6-332a4ec426c1",
									Text:        "Which of the following best describes 'Semantic HTML'?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "612d51e2-12a9-4cbd-b83d-f089912e8de9",
											Text:        "HTML that boosts the website's SEO",
											IsCorrect:   false,
											Description: "While semantic HTML can positively influence SEO, its primary purpose isn't solely for boosting SEO.",
										},
										{
											ID:          "b98d4e15-4ded-40b9-9649-ac1ec77b8273",
											Text:        "HTML with embedded multimedia elements",
											IsCorrect:   false,
											Description: "Embedding multimedia elements doesn't make HTML semantic.",
										},
										{
											ID:          "adadef45-8699-4be5-a2c9-719a4305799b",
											Text:        "HTML where elements are used according to their meaning, not just their appearance",
											IsCorrect:   true,
											Description: "Semantic HTML involves using HTML elements for their given meaning, which helps in improving web accessibility, making websites more readable for both users and machines.",
										},
										{
											ID:          "b31af3dc-83b1-4fd1-96cc-a94e404d1f74",
											Text:        "HTML that includes animations and dynamic content",
											IsCorrect:   false,
											Description: "Animations and dynamic content are unrelated to the semantics of an HTML document.",
										},
									},
								},
								{
									ID:          "483d743d-4b39-422d-a391-099bbcb47a63",
									Text:        "Which property would you use in CSS if you wanted to set the space between the content of an element and its border?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "b32e3ac5-4191-4119-af22-20110f972c4f",
											Text:        "margin",
											IsCorrect:   false,
											Description: "margin is the space outside the border, between the element and its surrounding elements.",
										},
										{
											ID:          "bb49bc81-ab30-437b-8d84-fa6554d3ff37",
											Text:        "border",
											IsCorrect:   false,
											Description: "border sets the border around an element but doesn't deal with spacing.",
										},
										{
											ID:          "d60a48f3-1816-4df8-9a57-dbaac5dc3aef",
											Text:        "spacing",
											IsCorrect:   false,
											Description: "spacing is not a standard CSS property.",
										},
										{
											ID:          "ecc3da18-8889-4b79-b5b5-ed4a5e957b28",
											Text:        "padding",
											IsCorrect:   true,
											Description: "Padding in CSS refers to the space between the content of an element and its border. It's an inner spacing.",
										},
									},
								},
								{
									ID:          "8636ca4c-26d8-4061-b135-7a5d589e66ec",
									Text:        "Which of the following pseudo-classes targets elements based on their position in a parent element?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "158f45d1-4ece-458e-886f-2edbfd47f0ff",
											Text:        ":first-child",
											IsCorrect:   true,
											Description: "The :first-child pseudo-class targets the first child element of a parent.",
										},
										{
											ID:          "19a6617a-6442-451d-8843-6a250ecaaacd",
											Text:        ":hover",
											IsCorrect:   false,
											Description: ":hover targets an element when it's being hovered over.",
										},
										{
											ID:          "c52aac43-90d0-4b36-9e67-4498cf2df5a0",
											Text:        ":active",
											IsCorrect:   false,
											Description: ":active targets an element, like a button, during the active state (e.g., when it's pressed).",
										},
										{
											ID:          "1e3ba4b1-104d-498c-9c2a-4384682d0b9c",
											Text:        ":visited",
											IsCorrect:   false,
											Description: ":visited targets links that have been visited.",
										},
									},
								},
								{
									ID:          "232d1d67-d33c-4047-b3d2-ab35b56f67d0",
									Text:        "Which HTML5 element is specifically designed to contain navigation links?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "65fd438a-a6b3-496a-b3e5-d6b9c7e2af49",
											Text:        "<sidebar>",
											IsCorrect:   false,
											Description: "There is no standard <sidebar> element in HTML.",
										},
										{
											ID:          "b1f846d4-d7e0-4690-9bd2-8910c59253e9",
											Text:        "<navbar>",
											IsCorrect:   false,
											Description: "While navbar is commonly used as a term in certain frameworks (e.g., Bootstrap), there's no <navbar> tag in standard HTML.",
										},
										{
											ID:          "b0609bd8-3e1b-4aab-bcc7-96cdf3eccd3f",
											Text:        "<menu>",
											IsCorrect:   false,
											Description: "<menu> isn't specifically for navigation links.",
										},
										{
											ID:          "ec10dbd4-f889-425c-b1bd-d7ed2524aecd",
											Text:        "<nav>",
											IsCorrect:   true,
											Description: "The <nav> element in HTML5 is specifically meant to enclose navigation links.",
										},
									},
								},
								{
									ID:          "f3844648-6b7f-443c-b146-da5e1c41376d",
									Text:        "What does the 'Box Model' in CSS refer to?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "22bdfe89-92ae-40fb-88eb-3271a8be0592",
											Text:        "A technique for 3D transformations.",
											IsCorrect:   false,
											Description: "3D transformations in CSS are achieved using different techniques, not the Box Model.",
										},
										{
											ID:          "e06f1697-6562-430c-ac61-b4a9d3512867",
											Text:        "The packaging of CSS files into boxes.",
											IsCorrect:   false,
											Description: "CSS files aren't packaged into boxes.",
										},
										{
											ID:          "f358267e-f693-4ca4-88f2-fe6671fff610",
											Text:        "A combination of padding, border, margin, and the actual content.",
											IsCorrect:   true,
											Description: "The CSS Box Model describes the design and layout by placing elements in a box with specific properties like padding, border, and margin.",
										},
										{
											ID:          "271a1d2c-de58-4671-b086-717d897398b5",
											Text:        "The grid system used in modern layouts.",
											IsCorrect:   false,
											Description: "While the grid system is vital for layouts, the Box Model specifically deals with the design and layout of individual elements.",
										},
									},
								},
								{
									ID:          "9fb2aea8-f16a-48b7-ae8b-bd909fd5e3f4",
									Text:        "In a flexbox container, which property is used to align the items vertically (assuming a row-direction)?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											ID:          "fc6658ec-f545-4468-82e0-9f7943acc724",
											Text:        "align-horizontal",
											IsCorrect:   false,
											Description: "`align-horizontal` is not a standard CSS property.",
										},
										{
											ID:          "7c91a113-9540-48c2-9812-e105067b14aa",
											Text:        "align-items",
											IsCorrect:   true,
											Description: "In a flexbox container, when you want to align items vertically in the case of a row direction, you'd use `align-items`.",
										},
										{
											ID:          "4f9f5dc4-6edd-4f0b-8620-857394ff543b",
											Text:        "justify-content",
											IsCorrect:   false,
											Description: "`justify-content` is used to align flex items along the main axis (horizontally in the case of a row direction).",
										},
										{
											ID:          "c0822cea-18f7-41ba-bc80-6d6f95b6ea0b",
											Text:        "vertical-align",
											IsCorrect:   false,
											Description: "`vertical-align` is used with inline-level and table-cell elements, not with flex items.",
										},
									},
								},
							},
						},
					},
				},
			},
		}, nil
	}

	return course.Course{}, course.ErrCourseNotFound
}
