using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Nordlys.Data;

namespace Nordlys.Controllers;

/// <summary>
/// Currency controller
/// </summary>
[ApiController]
[Route("api/currencies")]
public class CurrencyController : ControllerBase
{
    private readonly ILogger<CurrencyController> _logger;
    private readonly ApplicationDbContext _applicationDb;

    public CurrencyController(ApplicationDbContext applicationDb,
        ILogger<CurrencyController> logger)
    {
        _applicationDb = applicationDb;
        _logger = logger;
    }

    /// <summary>
    /// Retrieve all registered currencies.
    /// </summary>
    [HttpGet]
    public async Task<IEnumerable<Currency>> GetAll()
    {
        _logger.LogDebug("Called GetAll endpoint.");
        return await _applicationDb.Currencies.ToListAsync();
    }

    /// <summary>
    /// Retrieve a currency by its id.
    /// </summary>
    [HttpGet]
    [Route("{id}")]
    public async Task<IActionResult> GetById(int id)
    {
        _logger.LogDebug($"Called GetById with id : {id}");
        var result = await _applicationDb.Currencies.FindAsync(id);
        
        if (result == null)
        {
            return NotFound();
        }
        
        return Ok(result);
    }

    /// <summary>
    /// Delete a currency by its id.
    /// </summary>
    /// <param name="id">Currency ID</param>
    [HttpDelete]
    [Route("{id}")]
    public async Task<IActionResult> Delete(int id)
    {
        _logger.LogDebug($"Called Delete with id {id}");
        var entity = await _applicationDb.Currencies.FindAsync(id);
        if (entity == null)
        {
            return NotFound();
        }

        _applicationDb.Currencies.Remove(entity);
        return Ok();
    }
}