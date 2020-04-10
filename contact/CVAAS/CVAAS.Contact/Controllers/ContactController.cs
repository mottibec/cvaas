using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace CVAAS.Contact.Controllers
{
    [ApiController, Route("[controller]")]
    public class ContactController : ControllerBase
    {
        private readonly ILogger<ContactController> logger;

        public ContactController()
        {

        }

        [HttpGet]
        public async Task<IActionResult> Get()
        {

        }

    }
}