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

const CourseID = "fcf7890f-9c72-46d3-931e-34494307be37"
const CourseStepID = "c7486278-a50c-4629-89b9-cc1c74d7a538"
const QuizID = "fcf7890f-9c72-46d3-931e-34494307be37"

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
									Text:        "What is HTML used for?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "Styling the website",
											IsCorrect:   false,
											Description: "Styling a website is primarily done using CSS, which helps in determining colors, font sizes, positioning, and other visual elements.",
										},
										{
											Text:        "Describing the structure of a webpage",
											IsCorrect:   true,
											Description: "HTML stands for HyperText Markup Language and it provides the basic structure to a webpage. Through HTML, the browser understands headings, paragraphs, links, and other content types.",
										},
										{
											Text:        "Creating animations for a webpage",
											IsCorrect:   false,
											Description: "While animations can be added to a webpage, they are typically achieved using both CSS and JavaScript, depending on the complexity. HTML itself isn't used for animations.",
										},
										{
											Text:        "Accessing web databases",
											IsCorrect:   false,
											Description: "Accessing databases is generally the realm of backend languages and APIs, not directly through HTML.",
										},
									},
								},
								{
									Text:        "Which of the following is NOT an HTML element?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "<heading>",
											IsCorrect:   false,
											Description: "The term <heading> is not used as an HTML element. Instead, we have heading tags that range from <h1> to <h6> to denote headings of different levels.",
										},
										{
											Text:        "<paragraph>",
											IsCorrect:   false,
											Description: "<paragraph> is not the correct HTML element for creating paragraphs. The right tag for a paragraph in HTML is <p>.",
										},
										{
											Text:        "<list>",
											IsCorrect:   false,
											Description: "While you can create lists in HTML, the tag <list> doesn't exist. For lists, we use <ul> for unordered lists and <ol> for ordered lists.",
										},
										{
											Text:        "None of the above",
											IsCorrect:   true,
											Description: "All of the provided options are not standard HTML elements.",
										},
									},
								},
								{
									Text:        "Which language is primarily used to determine the visual style of a website?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "JavaScript",
											IsCorrect:   false,
											Description: "JavaScript is mainly used for adding functionality and interactivity to web pages.",
										},
										{
											Text:        "Python",
											IsCorrect:   false,
											Description: "Python is a programming language that's often used for backend web development, data analysis, artificial intelligence, and more.",
										},
										{
											Text:        "HTML",
											IsCorrect:   false,
											Description: "HTML is the markup language used to structure content on the web. While it can carry some style attributes, CSS is the primary tool for styling.",
										},
										{
											Text:        "CSS",
											IsCorrect:   true,
											Description: "CSS (Cascading Style Sheets) is the language specifically designed to style web content. It defines how elements should appear in terms of layout, colors, fonts, and more.",
										},
									},
								},
								{
									Text:        "What is the primary purpose of visual design in web development?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "Accelerating website speed",
											IsCorrect:   false,
											Description: "While a well-designed website can improve user experience, the speed of a website is more about optimization, efficient coding, and hosting solutions.",
										},
										{
											Text:        "Ensuring website security",
											IsCorrect:   false,
											Description: "Website security is crucial, but visual design isn't directly related to it. Security measures involve both frontend and backend strategies, encryptions, and more.",
										},
										{
											Text:        "Delivering a unique message and enhancing user experience",
											IsCorrect:   true,
											Description: "Visual design uses typography, color, graphics, and more to help convey a unique message and provide a memorable user experience.",
										},
										{
											Text:        "Backing up website data",
											IsCorrect:   false,
											Description: "Backing up website data is an administrative task and not a direct responsibility of visual design.",
										},
									},
								},
								{
									Text:        "Why is web accessibility important?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "It ensures that websites look the same on all devices.",
											IsCorrect:   false,
											Description: "While a consistent appearance across devices is essential, that's more about responsive design than accessibility.",
										},
										{
											Text:        "It makes websites load faster.",
											IsCorrect:   false,
											Description: "Web accessibility doesn't necessarily mean faster load times. Load times depend on optimization techniques, server speeds, and more.",
										},
										{
											Text:        "It guarantees that websites can be understood and navigated by everyone, including people with disabilities.",
											IsCorrect:   true,
											Description: "Accessibility ensures that web content and user interfaces are accessible to all users, regardless of physical or cognitive disabilities, ensuring inclusivity on the web.",
										},
										{
											Text:        "It ensures higher search engine rankings.",
											IsCorrect:   false,
											Description: "While search engines do value accessible websites, the primary goal of web accessibility is inclusiveness, not SEO.",
										},
									},
								},
								{
									Text:        "What does responsive web design ensure?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "Faster loading speed on mobile devices",
											IsCorrect:   false,
											Description: "Responsive design doesn't inherently speed up website loading times. Its main goal is adaptability.",
										},
										{
											Text:        "Higher security against web threats",
											IsCorrect:   false,
											Description: "Responsive design focuses on layout adaptability, not security.",
										},
										{
											Text:        "Flexibility in website appearance according to device size and orientation",
											IsCorrect:   true,
											Description: "Responsive web design allows websites to adjust and look good on devices of all sizes, whether it's a large monitor, tablet, or smartphone.",
										},
										{
											Text:        "Vibrant color schemes for websites",
											IsCorrect:   false,
											Description: "Color schemes are part of visual design. Responsiveness is about layout adjustments based on screen sizes.",
										},
									},
								},
								{
									Text:        "Which of the following is a powerful layout method introduced in CSS3?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "CSS Shapes",
											IsCorrect:   false,
											Description: "CSS Shapes is about wrapping content around custom paths, not about layout per se.",
										},
										{
											Text:        "CSS Flexbox",
											IsCorrect:   true,
											Description: "CSS Flexbox provides an efficient way to lay out, align, and distribute space among items in a container, even when their sizes are unknown or dynamic.",
										},
										{
											Text:        "CSS Templates",
											IsCorrect:   false,
											Description: "There's no standard feature named 'CSS Templates'.",
										},
										{
											Text:        "CSS Fonts",
											IsCorrect:   false,
											Description: "CSS Fonts relates to font-face and loading custom fonts, not layouts.",
										},
									},
								},
								{
									Text:        "What does the CSS Grid enable in terms of layout design?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "It allows embedding multimedia.",
											IsCorrect:   false,
											Description: "While multimedia can be placed within a CSS Grid, the grid itself is not specifically for embedding multimedia.",
										},
										{
											Text:        "It ensures website security.",
											IsCorrect:   false,
											Description: "CSS Grid focuses on layout and has nothing to do with security.",
										},
										{
											Text:        "It enables easy creation of flexible, two-dimensional layouts.",
											IsCorrect:   true,
											Description: "CSS Grid provides a method for placing elements within columns and rows, making complex responsive layouts more manageable.",
										},
										{
											Text:        "It automates browser compatibility.",
											IsCorrect:   false,
											Description: "Browser compatibility depends on various factors. While Grid is supported in many modern browsers, its use doesn't 'automate' compatibility.",
										},
									},
								},
								{
									Text:        "Which of the following elements can describe text as a heading in HTML?",
									Difficulty:  "easy",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "<header>",
											IsCorrect:   false,
											Description: "The <header> element typically contains introductory content or navigation links and not used as a standard heading.",
										},
										{
											Text:        "<heading>",
											IsCorrect:   false,
											Description: "There is no standard <heading> tag in HTML.",
										},
										{
											Text:        "<h1>",
											IsCorrect:   true,
											Description: "The <h1> tag in HTML is used to represent the top-level heading, making it the most important among the heading tags ranging from <h1> to <h6>.",
										},
										{
											Text:        "<top>",
											IsCorrect:   false,
											Description: "The <top> tag doesn't exist in standard HTML.",
										},
									},
								},
								{
									Text:        "When you want to make a list of items without any specific order, which HTML tag would you use?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "<ol>",
											IsCorrect:   false,
											Description: "The <ol> tag is used for ordered lists, where the order of items is important.",
										},
										{
											Text:        "<dl>",
											IsCorrect:   false,
											Description: "The <dl> tag is for description lists and isn't typically used for general lists.",
										},
										{
											Text:        "<ul>",
											IsCorrect:   true,
											Description: "The <ul> tag stands for unordered list, which is used for lists where order doesn't matter.",
										},
										{
											Text:        "<list>",
											IsCorrect:   false,
											Description: "There is no standard <list> tag in HTML.",
										},
									},
								},
								{
									Text:        "Considering the 'Cascading' in Cascading Style Sheets (CSS), which of the following statements is true?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "External stylesheets have a higher priority than inline styles.",
											IsCorrect:   false,
											Description: "Inline styles have a higher priority than external stylesheets.",
										},
										{
											Text:        "The last rule defined overrides any previous, conflicting rules.",
											IsCorrect:   true,
											Description: "Cascading means that the order of CSS rules matters. If two rules conflict, the latter rule will take precedence, given they have the same specificity.",
										},
										{
											Text:        "All styles have an equal weight, regardless of where they are defined.",
											IsCorrect:   false,
											Description: "Styles do not have equal weight; their application is determined by specificity and order.",
										},
										{
											Text:        "Inline styles only apply to JavaScript-generated content.",
											IsCorrect:   false,
											Description: "Inline styles can be applied directly to any HTML element, not just JavaScript-generated content.",
										},
									},
								},
								{
									Text:        "In responsive design, what does the 'viewport' refer to?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "The visible area of a webpage that the user can interact with.",
											IsCorrect:   true,
											Description: "The viewport is the user's visible area of a web page. It varies with the device, and it can be controlled using the <meta> viewport element.",
										},
										{
											Text:        "The server's perspective on how to serve content.",
											IsCorrect:   false,
											Description: "The server doesn't have a 'perspective'; it serves content based on requests.",
										},
										{
											Text:        "The total area of a website, including off-screen content.",
											IsCorrect:   false,
											Description: "The viewport only concerns what's currently visible and does not include off-screen content.",
										},
										{
											Text:        "The framework used to build mobile-responsive designs.",
											IsCorrect:   false,
											Description: "While certain frameworks can help with responsive design, the term 'viewport' does not pertain to any specific framework.",
										},
									},
								},
								{
									Text:        "Given the growth of mobile web browsing, why is Flexbox considered an important tool in responsive design?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "It enables developers to use less JavaScript.",
											IsCorrect:   false,
											Description: "While Flexbox can reduce the need for some layout-related JavaScript, its main purpose isn't about reducing JavaScript usage.",
										},
										{
											Text:        "It allows more vibrant color schemes.",
											IsCorrect:   false,
											Description: "Flexbox is about layout, not color schemes.",
										},
										{
											Text:        "Flexbox layouts adapt to different screen sizes without requiring pixel-based dimensions.",
											IsCorrect:   true,
											Description: "Flexbox allows for flexible layouts that can adapt to various screen sizes and orientations, making it easier to create responsive designs without relying heavily on fixed dimensions.",
										},
										{
											Text:        "It's a plugin that makes websites mobile-friendly automatically.",
											IsCorrect:   false,
											Description: "Flexbox is not a plugin; it's a CSS layout model.",
										},
									},
								},
								{
									Text:        "In the context of CSS Grid, what does the 'fr' unit stand for?",
									Difficulty:  "medium",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "Frame",
											IsCorrect:   false,
											Description: "While 'frame' sounds related to grids or layouts, it's not what 'fr' represents.",
										},
										{
											Text:        "Fraction",
											IsCorrect:   true,
											Description: "The 'fr' unit in CSS Grid stands for 'fraction'. It represents a fraction of the available space in the grid container.",
										},
										{
											Text:        "Frequency",
											IsCorrect:   false,
											Description: "Frequency is not a unit of measurement in CSS Grid.",
										},
										{
											Text:        "Format",
											IsCorrect:   false,
											Description: "Format doesn't correlate with the purpose of 'fr' in grid layouts.",
										},
									},
								},
								{
									Text:        "Which of the following best describes 'Semantic HTML'?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "HTML that boosts the website's SEO",
											IsCorrect:   false,
											Description: "While semantic HTML can positively influence SEO, its primary purpose isn't solely for boosting SEO.",
										},
										{
											Text:        "HTML with embedded multimedia elements",
											IsCorrect:   false,
											Description: "Embedding multimedia elements doesn't make HTML semantic.",
										},
										{
											Text:        "HTML where elements are used according to their meaning, not just their appearance",
											IsCorrect:   true,
											Description: "Semantic HTML involves using HTML elements for their given meaning, which helps in improving web accessibility, making websites more readable for both users and machines.",
										},
										{
											Text:        "HTML that includes animations and dynamic content",
											IsCorrect:   false,
											Description: "Animations and dynamic content are unrelated to the semantics of an HTML document.",
										},
									},
								},
								{
									Text:        "Which property would you use in CSS if you wanted to set the space between the content of an element and its border?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "margin",
											IsCorrect:   false,
											Description: "margin is the space outside the border, between the element and its surrounding elements.",
										},
										{
											Text:        "border",
											IsCorrect:   false,
											Description: "border sets the border around an element but doesn't deal with spacing.",
										},
										{
											Text:        "spacing",
											IsCorrect:   false,
											Description: "spacing is not a standard CSS property.",
										},
										{
											Text:        "padding",
											IsCorrect:   true,
											Description: "Padding in CSS refers to the space between the content of an element and its border. It's an inner spacing.",
										},
									},
								},
								{
									Text:        "Which of the following pseudo-classes targets elements based on their position in a parent element?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        ":first-child",
											IsCorrect:   true,
											Description: "The :first-child pseudo-class targets the first child element of a parent.",
										},
										{
											Text:        ":hover",
											IsCorrect:   false,
											Description: ":hover targets an element when it's being hovered over.",
										},
										{
											Text:        ":active",
											IsCorrect:   false,
											Description: ":active targets an element, like a button, during the active state (e.g., when it's pressed).",
										},
										{
											Text:        ":visited",
											IsCorrect:   false,
											Description: ":visited targets links that have been visited.",
										},
									},
								},
								{
									Text:        "Which HTML5 element is specifically designed to contain navigation links?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "<sidebar>",
											IsCorrect:   false,
											Description: "There is no standard <sidebar> element in HTML.",
										},
										{
											Text:        "<navbar>",
											IsCorrect:   false,
											Description: "While navbar is commonly used as a term in certain frameworks (e.g., Bootstrap), there's no <navbar> tag in standard HTML.",
										},
										{
											Text:        "<menu>",
											IsCorrect:   false,
											Description: "<menu> isn't specifically for navigation links.",
										},
										{
											Text:        "<nav>",
											IsCorrect:   true,
											Description: "The <nav> element in HTML5 is specifically meant to enclose navigation links.",
										},
									},
								},
								{
									Text:        "What does the 'Box Model' in CSS refer to?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "A technique for 3D transformations.",
											IsCorrect:   false,
											Description: "3D transformations in CSS are achieved using different techniques, not the Box Model.",
										},
										{
											Text:        "The packaging of CSS files into boxes.",
											IsCorrect:   false,
											Description: "CSS files aren't packaged into boxes.",
										},
										{
											Text:        "A combination of padding, border, margin, and the actual content.",
											IsCorrect:   true,
											Description: "The CSS Box Model describes the design and layout by placing elements in a box with specific properties like padding, border, and margin.",
										},
										{
											Text:        "The grid system used in modern layouts.",
											IsCorrect:   false,
											Description: "While the grid system is vital for layouts, the Box Model specifically deals with the design and layout of individual elements.",
										},
									},
								},
								{
									Text:        "In a flexbox container, which property is used to align the items vertically (assuming a row-direction)?",
									Difficulty:  "hard",
									Rating:      4.3,
									RatingCount: 1991,
									Answers: []course.QuizAnswer{
										{
											Text:        "align-horizontal",
											IsCorrect:   false,
											Description: "`align-horizontal` is not a standard CSS property.",
										},
										{
											Text:        "align-items",
											IsCorrect:   true,
											Description: "In a flexbox container, when you want to align items vertically in the case of a row direction, you'd use `align-items`.",
										},
										{
											Text:        "justify-content",
											IsCorrect:   false,
											Description: "`justify-content` is used to align flex items along the main axis (horizontally in the case of a row direction).",
										},
										{
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
